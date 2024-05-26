package app

import (
	artistService "github.com/rauzh/cd-core/artist/service"
	managerService "github.com/rauzh/cd-core/manager/service"
	publicationService "github.com/rauzh/cd-core/publication/service"
	releaseService "github.com/rauzh/cd-core/release/service"
	"github.com/rauzh/cd-core/reporter/service"
	"github.com/rauzh/cd-core/requests/base"
	requestService "github.com/rauzh/cd-core/requests/base/service"
	statService "github.com/rauzh/cd-core/statistics/service"
	trackService "github.com/rauzh/cd-core/track/service"
	userService "github.com/rauzh/cd-core/user/service"
)

type App struct {
	UseCases *AppUseCases
	Services *AppServices
}

type AppUseCases struct {
	SignContractReqUC base.IRequestUseCase
	PublishReqUC      base.IRequestUseCase
}

type AppServices struct {
	ArtistService      artistService.IArtistService
	ManagerService     managerService.IManagerService
	PublicationService publicationService.IPublicationService
	ReleaseService     releaseService.IReleaseService
	ReportService      service.IReportService
	RequestService     requestService.IRequestService
	StatService        statService.IStatisticsService
	TrackService       trackService.ITrackService
	UserService        userService.IUserService
}
