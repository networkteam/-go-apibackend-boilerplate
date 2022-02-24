package finder

import (
	"context"

	"myvendor.mytld/myproject/backend/domain"
	"myvendor.mytld/myproject/backend/persistence/repository"
	"myvendor.mytld/myproject/backend/security/authentication"
	"myvendor.mytld/myproject/backend/security/authorization"
)

func (f *Finder) QueryOrganisation(ctx context.Context, query domain.OrganisationQuery) (domain.Organisation, error) {
	authorizer := authorization.NewAuthorizer(authentication.GetAuthContext(ctx))
	err := authorizer.AllowsOrganisationQuery(query)
	if err != nil {
		return domain.Organisation{}, err
	}

	record, err := repository.FindOrganisationByID(ctx, f.db, query.OrganisationID)
	if err != nil {
		return record, err
	}
	return record, nil
}

func (f *Finder) QueryOrganisations(ctx context.Context, query domain.OrganisationsQuery, paging repository.Paging) ([]domain.Organisation, error) {
	authorizer := authorization.NewAuthorizer(authentication.GetAuthContext(ctx))
	err := authorizer.AllowsAndFilterAllOrganisationsQuery(&query)
	if err != nil {
		return nil, err
	}
	return repository.FindAllOrganisations(ctx, f.db, paging, query)
}

func (f *Finder) CountOrganisations(ctx context.Context, query domain.OrganisationsQuery) (int, error) {
	authorizer := authorization.NewAuthorizer(authentication.GetAuthContext(ctx))
	err := authorizer.AllowsAndFilterAllOrganisationsQuery(&query)
	if err != nil {
		return 0, err
	}

	return repository.CountAllOrganisations(ctx, f.db, query)
}
