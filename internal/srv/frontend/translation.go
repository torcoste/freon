package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/filter"

	"github.com/AlekSi/pointer"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

func (srv *server) createTranslation(params op.CreateTranslationParams, session *app.UserSession) op.CreateTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.CreateTranslation(
		ctx,
		session.UserID,
		swag.Int64Value(params.Args.LocalizationID),
		swag.Int64Value(params.Args.IdentifierID),
		swag.StringValue(params.Args.Singular),
		params.Args.Plural,
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errCreateTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewCreateTranslationNoContent()
}

func (srv *server) listTranslations(params op.ListTranslationsParams, session *app.UserSession) op.ListTranslationsResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	entities, err := srv.app.GetTranslations(
		ctx,
		filter.TranslationFilter{
			LocalizationID: swag.Int64Value(params.LocalizationID),
		},
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errListTranslations(log, err, codeInternal)
	case nil:
	}

	return op.NewListTranslationsOK().WithPayload(apiArrayTranslation(entities))
}

func (srv *server) deleteTranslation(params op.DeleteTranslationParams, session *app.UserSession) op.DeleteTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.DeleteTranslation(ctx, params.ID)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errDeleteTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewDeleteTranslationNoContent()
}

func (srv *server) updateTranslation(params op.UpdateTranslationParams, session *app.UserSession) op.UpdateTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.UpdateTranslation(
		ctx,
		params.ID,
		swag.StringValue(params.Args.Singular),
		params.Args.Plural,
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errUpdateTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewUpdateTranslationNoContent()
}

func (srv *server) statusTranslation(params op.StatusTranslationParams, session *app.UserSession) op.StatusTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.UpdateStatusTranslation(ctx, params.ID, params.Status)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errStatusTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewStatusTranslationNoContent()
}

func (srv *server) autoTranslation(params op.AutoTranslationParams, session *app.UserSession) op.AutoTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	source, err := language.Parse(swag.StringValue(params.Args.Source))
	if err != nil {
		return errAutoTranslation(log, err, codeInternal)
	}

	target, err := language.Parse(swag.StringValue(params.Args.Target))
	if err != nil {
		return errAutoTranslation(log, err, codeInternal)
	}

	translate, err := srv.app.Translate(ctx, swag.StringValue(params.Args.Text), source, target)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errAutoTranslation(log, err, codeInternal)
	case app.ErrAutoTranslation:
		return errAutoTranslation(log, err, codeAutoTranslationNotSupported)
	case nil:
	}

	return op.NewAutoTranslationOK().WithPayload(&op.AutoTranslationOKBody{
		Text: pointer.ToString(translate),
	})
}
