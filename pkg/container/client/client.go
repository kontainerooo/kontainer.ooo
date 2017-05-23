package client

import (
  grpctransport "github.com/go-kit/kit/transport/grpc"
  "github.com/go-kit/kit/log"
  "github.com/go-kit/kit/endpoint"

  "github.com/kontainerooo/kontainer.ooo/pkg/pb"
  "github.com/kontainerooo/kontainer.ooo/pkg/container"
)


// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *container.Endpoints {
  
    var CreateContainerEndpoint endpoint.Endpoint
			{
				CreateContainerEndpoint = grpctransport.NewClient(
					conn,
					"containerService",
					"CreateContainer",
					EncodeGRPCCreateContainerRequest,
					DecodeGRPCCreateContainerResponse,
					pb.CreateContainerResponse{},
				).Endpoint()
			}
  
    var RemoveContainerEndpoint endpoint.Endpoint
			{
				RemoveContainerEndpoint = grpctransport.NewClient(
					conn,
					"containerService",
					"RemoveContainer",
					EncodeGRPCRemoveContainerRequest,
					DecodeGRPCRemoveContainerResponse,
					pb.RemoveContainerResponse{},
				).Endpoint()
			}
  
    var InstancesEndpoint endpoint.Endpoint
			{
				InstancesEndpoint = grpctransport.NewClient(
					conn,
					"containerService",
					"Instances",
					EncodeGRPCInstancesRequest,
					DecodeGRPCInstancesResponse,
					pb.InstancesResponse{},
				).Endpoint()
			}
  
    var StartContainerEndpoint endpoint.Endpoint
			{
				StartContainerEndpoint = grpctransport.NewClient(
					conn,
					"containerService",
					"StartContainer",
					EncodeGRPCStartContainerRequest,
					DecodeGRPCStartContainerResponse,
					pb.StartContainerResponse{},
				).Endpoint()
			}
  
    var StopContainerEndpoint endpoint.Endpoint
			{
				StopContainerEndpoint = grpctransport.NewClient(
					conn,
					"containerService",
					"StopContainer",
					EncodeGRPCStopContainerRequest,
					DecodeGRPCStopContainerResponse,
					pb.StopContainerResponse{},
				).Endpoint()
			}
  
    var IsRunningEndpoint endpoint.Endpoint
			{
				IsRunningEndpoint = grpctransport.NewClient(
					conn,
					"containerService",
					"IsRunning",
					EncodeGRPCIsRunningRequest,
					DecodeGRPCIsRunningResponse,
					pb.IsRunningResponse{},
				).Endpoint()
			}
  
    var ExecuteEndpoint endpoint.Endpoint
			{
				ExecuteEndpoint = grpctransport.NewClient(
					conn,
					"containerService",
					"Execute",
					EncodeGRPCExecuteRequest,
					DecodeGRPCExecuteResponse,
					pb.ExecuteResponse{},
				).Endpoint()
			}
  

        return &container.Endpoints { 
          CreateContainerEndpoint: CreateContainerEndpoint,
          RemoveContainerEndpoint: RemoveContainerEndpoint,
          InstancesEndpoint: InstancesEndpoint,
          StartContainerEndpoint: StartContainerEndpoint,
          StopContainerEndpoint: StopContainerEndpoint,
          IsRunningEndpoint: IsRunningEndpoint,
          ExecuteEndpoint: ExecuteEndpoint,
        }
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}


  // EncodeGRPCCreateContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
  // messages/container.proto-domain createcontainer request to a gRPC CreateContainer request.
	func EncodeGRPCCreateContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*container.CreateContainerRequest)
		return &pb.CreateContainerRequest{

    }, nil
	}

  // DecodeGRPCCreateContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
  // gRPC CreateContainer response to a messages/container.proto-domain createcontainer response.
	func DecodeGRPCCreateContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
		response := grpcResponse.(*pb.CreateContainerResponse)
		return &container.CreateContainerResponse{
      Error: getError(response.Error),
    }, nil
	}

  // EncodeGRPCRemoveContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
  // messages/container.proto-domain removecontainer request to a gRPC RemoveContainer request.
	func EncodeGRPCRemoveContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*container.RemoveContainerRequest)
		return &pb.RemoveContainerRequest{

    }, nil
	}

  // DecodeGRPCRemoveContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
  // gRPC RemoveContainer response to a messages/container.proto-domain removecontainer response.
	func DecodeGRPCRemoveContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
		response := grpcResponse.(*pb.RemoveContainerResponse)
		return &container.RemoveContainerResponse{
      Error: getError(response.Error),
    }, nil
	}

  // EncodeGRPCInstancesRequest is a transport/grpc.EncodeRequestFunc that converts a
  // messages/container.proto-domain instances request to a gRPC Instances request.
	func EncodeGRPCInstancesRequest(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*container.InstancesRequest)
		return &pb.InstancesRequest{

    }, nil
	}

  // DecodeGRPCInstancesResponse is a transport/grpc.DecodeResponseFunc that converts a
  // gRPC Instances response to a messages/container.proto-domain instances response.
	func DecodeGRPCInstancesResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
		response := grpcResponse.(*pb.InstancesResponse)
		return &container.InstancesResponse{
      Error: getError(response.Error),
    }, nil
	}

  // EncodeGRPCStartContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
  // messages/container.proto-domain startcontainer request to a gRPC StartContainer request.
	func EncodeGRPCStartContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*container.StartContainerRequest)
		return &pb.StartContainerRequest{

    }, nil
	}

  // DecodeGRPCStartContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
  // gRPC StartContainer response to a messages/container.proto-domain startcontainer response.
	func DecodeGRPCStartContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
		response := grpcResponse.(*pb.StartContainerResponse)
		return &container.StartContainerResponse{
      Error: getError(response.Error),
    }, nil
	}

  // EncodeGRPCStopContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
  // messages/container.proto-domain stopcontainer request to a gRPC StopContainer request.
	func EncodeGRPCStopContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*container.StopContainerRequest)
		return &pb.StopContainerRequest{

    }, nil
	}

  // DecodeGRPCStopContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
  // gRPC StopContainer response to a messages/container.proto-domain stopcontainer response.
	func DecodeGRPCStopContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
		response := grpcResponse.(*pb.StopContainerResponse)
		return &container.StopContainerResponse{
      Error: getError(response.Error),
    }, nil
	}

  // EncodeGRPCIsRunningRequest is a transport/grpc.EncodeRequestFunc that converts a
  // messages/container.proto-domain isrunning request to a gRPC IsRunning request.
	func EncodeGRPCIsRunningRequest(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*container.IsRunningRequest)
		return &pb.IsRunningRequest{

    }, nil
	}

  // DecodeGRPCIsRunningResponse is a transport/grpc.DecodeResponseFunc that converts a
  // gRPC IsRunning response to a messages/container.proto-domain isrunning response.
	func DecodeGRPCIsRunningResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
		response := grpcResponse.(*pb.IsRunningResponse)
		return &container.IsRunningResponse{
      Error: getError(response.Error),
    }, nil
	}

  // EncodeGRPCExecuteRequest is a transport/grpc.EncodeRequestFunc that converts a
  // messages/container.proto-domain execute request to a gRPC Execute request.
	func EncodeGRPCExecuteRequest(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*container.ExecuteRequest)
		return &pb.ExecuteRequest{

    }, nil
	}

  // DecodeGRPCExecuteResponse is a transport/grpc.DecodeResponseFunc that converts a
  // gRPC Execute response to a messages/container.proto-domain execute response.
	func DecodeGRPCExecuteResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
		response := grpcResponse.(*pb.ExecuteResponse)
		return &container.ExecuteResponse{
      Error: getError(response.Error),
    }, nil
	}

