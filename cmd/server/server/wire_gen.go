// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/authgear/authgear-nft-indexer/pkg/handler"
	"github.com/authgear/authgear-nft-indexer/pkg/mutator"
	"github.com/authgear/authgear-nft-indexer/pkg/query"
	"github.com/authgear/authgear-nft-indexer/pkg/web3"
	"github.com/authgear/authgear-server/pkg/util/httputil"
	"net/http"
)

// Injectors from wire.go:

func NewHealthCheckAPIHandler(p *handler.RequestProvider) http.Handler {
	factory := p.LogFactory
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	registerCollectionHandlerLogger := handler.NewRegisterCollectionHandlerLogger(factory)
	config := p.Config
	db := p.Database
	request := p.Request
	context := handler.ProvideRequestContext(request)
	healthCheckAPIHandler := &handler.HealthCheckAPIHandler{
		JSON:     jsonResponseWriter,
		Logger:   registerCollectionHandlerLogger,
		Config:   config,
		Database: db,
		Context:  context,
	}
	return healthCheckAPIHandler
}

func NewRegisterCollectionAPIHandler(p *handler.RequestProvider) http.Handler {
	factory := p.LogFactory
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	registerCollectionHandlerLogger := handler.NewRegisterCollectionHandlerLogger(factory)
	config := p.Config
	alchemyAPI := &web3.AlchemyAPI{
		Config: config,
	}
	request := p.Request
	context := handler.ProvideRequestContext(request)
	db := p.Database
	nftCollectionMutator := &mutator.NFTCollectionMutator{
		Ctx:     context,
		Session: db,
	}
	registerCollectionAPIHandler := &handler.RegisterCollectionAPIHandler{
		JSON:                 jsonResponseWriter,
		Logger:               registerCollectionHandlerLogger,
		Config:               config,
		AlchemyAPI:           alchemyAPI,
		NFTCollectionMutator: nftCollectionMutator,
	}
	return registerCollectionAPIHandler
}

func NewDeregisterCollectionAPIHandler(p *handler.RequestProvider) http.Handler {
	factory := p.LogFactory
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	registerCollectionHandlerLogger := handler.NewRegisterCollectionHandlerLogger(factory)
	config := p.Config
	request := p.Request
	context := handler.ProvideRequestContext(request)
	db := p.Database
	nftCollectionMutator := &mutator.NFTCollectionMutator{
		Ctx:     context,
		Session: db,
	}
	deregisterCollectionAPIHandler := &handler.DeregisterCollectionAPIHandler{
		JSON:                 jsonResponseWriter,
		Logger:               registerCollectionHandlerLogger,
		Config:               config,
		NFTCollectionMutator: nftCollectionMutator,
	}
	return deregisterCollectionAPIHandler
}

func NewListCollectionAPIHandler(p *handler.RequestProvider) http.Handler {
	factory := p.LogFactory
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	listCollectionHandlerLogger := handler.NewListCollectionHandlerLogger(factory)
	request := p.Request
	context := handler.ProvideRequestContext(request)
	db := p.Database
	nftCollectionQuery := &query.NFTCollectionQuery{
		Ctx:     context,
		Session: db,
	}
	listCollectionAPIHandler := &handler.ListCollectionAPIHandler{
		JSON:               jsonResponseWriter,
		Logger:             listCollectionHandlerLogger,
		NFTCollectionQuery: nftCollectionQuery,
	}
	return listCollectionAPIHandler
}

func NewListCollectionOwnerAPIHandler(p *handler.RequestProvider) http.Handler {
	factory := p.LogFactory
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	listCollectionOwnerHandlerLogger := handler.NewListCollectionOwnerHandlerLogger(factory)
	request := p.Request
	context := handler.ProvideRequestContext(request)
	db := p.Database
	nftOwnerQuery := query.NFTOwnerQuery{
		Ctx:     context,
		Session: db,
	}
	listCollectionOwnersAPIHandler := &handler.ListCollectionOwnersAPIHandler{
		JSON:          jsonResponseWriter,
		Logger:        listCollectionOwnerHandlerLogger,
		NFTOwnerQuery: nftOwnerQuery,
	}
	return listCollectionOwnersAPIHandler
}

func NewListOwnerNFTAPIHandler(p *handler.RequestProvider) http.Handler {
	factory := p.LogFactory
	jsonResponseWriterLogger := httputil.NewJSONResponseWriterLogger(factory)
	jsonResponseWriter := &httputil.JSONResponseWriter{
		Logger: jsonResponseWriterLogger,
	}
	listOwnerNFTHandlerLogger := handler.NewListOwnerNFTHandlerLogger(factory)
	config := p.Config
	request := p.Request
	context := handler.ProvideRequestContext(request)
	db := p.Database
	nftOwnerQuery := query.NFTOwnerQuery{
		Ctx:     context,
		Session: db,
	}
	listOwnerNFTAPIHandler := &handler.ListOwnerNFTAPIHandler{
		JSON:          jsonResponseWriter,
		Logger:        listOwnerNFTHandlerLogger,
		Config:        config,
		NFTOwnerQuery: nftOwnerQuery,
	}
	return listOwnerNFTAPIHandler
}
