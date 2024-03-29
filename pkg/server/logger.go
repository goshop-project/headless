package server

import "darvaza.org/slog"

func (srv *Server) fatal(err error) slog.Logger {
	l := srv.cfg.Logger.Fatal()
	if err != nil {
		l = l.WithField(slog.ErrorFieldName, err)
	}
	return l
}

func (srv *Server) error(err error) slog.Logger {
	l := srv.cfg.Logger.Error()
	if err != nil {
		l = l.WithField(slog.ErrorFieldName, err)
	}
	return l
}

func (srv *Server) warn() slog.Logger {
	return srv.cfg.Logger.Warn()
}

func (srv *Server) info() slog.Logger {
	return srv.cfg.Logger.Info()
}

func (srv *Server) debug() slog.Logger {
	return srv.cfg.Logger.Debug()
}
