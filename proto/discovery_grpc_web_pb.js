/**
 * @fileoverview gRPC-Web generated client stub for discovery
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.discovery = require('./discovery_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.discovery.DiscoveryServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.discovery.DiscoveryServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.RegisterRequest,
 *   !proto.discovery.RegisterResponse>}
 */
const methodDescriptor_DiscoveryService_RegisterService = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryService/RegisterService',
  grpc.web.MethodType.UNARY,
  proto.discovery.RegisterRequest,
  proto.discovery.RegisterResponse,
  /**
   * @param {!proto.discovery.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.RegisterResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.RegisterRequest,
 *   !proto.discovery.RegisterResponse>}
 */
const methodInfo_DiscoveryService_RegisterService = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.RegisterResponse,
  /**
   * @param {!proto.discovery.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.RegisterResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.RegisterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.RegisterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceClient.prototype.registerService =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryService/RegisterService',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_RegisterService,
      callback);
};


/**
 * @param {!proto.discovery.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.RegisterResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServicePromiseClient.prototype.registerService =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryService/RegisterService',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_RegisterService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.DiscoverRequest,
 *   !proto.discovery.DiscoverResponse>}
 */
const methodDescriptor_DiscoveryService_Discover = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryService/Discover',
  grpc.web.MethodType.UNARY,
  proto.discovery.DiscoverRequest,
  proto.discovery.DiscoverResponse,
  /**
   * @param {!proto.discovery.DiscoverRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.DiscoverResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.DiscoverRequest,
 *   !proto.discovery.DiscoverResponse>}
 */
const methodInfo_DiscoveryService_Discover = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.DiscoverResponse,
  /**
   * @param {!proto.discovery.DiscoverRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.DiscoverResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.DiscoverRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.DiscoverResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.DiscoverResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceClient.prototype.discover =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryService/Discover',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_Discover,
      callback);
};


/**
 * @param {!proto.discovery.DiscoverRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.DiscoverResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServicePromiseClient.prototype.discover =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryService/Discover',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_Discover);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.ListRequest,
 *   !proto.discovery.ListResponse>}
 */
const methodDescriptor_DiscoveryService_ListAllServices = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryService/ListAllServices',
  grpc.web.MethodType.UNARY,
  proto.discovery.ListRequest,
  proto.discovery.ListResponse,
  /**
   * @param {!proto.discovery.ListRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.ListResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.ListRequest,
 *   !proto.discovery.ListResponse>}
 */
const methodInfo_DiscoveryService_ListAllServices = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.ListResponse,
  /**
   * @param {!proto.discovery.ListRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.ListResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.ListRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.ListResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.ListResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceClient.prototype.listAllServices =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryService/ListAllServices',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_ListAllServices,
      callback);
};


/**
 * @param {!proto.discovery.ListRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.ListResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServicePromiseClient.prototype.listAllServices =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryService/ListAllServices',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_ListAllServices);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.StateRequest,
 *   !proto.discovery.StateResponse>}
 */
const methodDescriptor_DiscoveryService_State = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryService/State',
  grpc.web.MethodType.UNARY,
  proto.discovery.StateRequest,
  proto.discovery.StateResponse,
  /**
   * @param {!proto.discovery.StateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.StateResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.StateRequest,
 *   !proto.discovery.StateResponse>}
 */
const methodInfo_DiscoveryService_State = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.StateResponse,
  /**
   * @param {!proto.discovery.StateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.StateResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.StateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.StateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.StateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceClient.prototype.state =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryService/State',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_State,
      callback);
};


/**
 * @param {!proto.discovery.StateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.StateResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServicePromiseClient.prototype.state =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryService/State',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_State);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.discovery.DiscoveryServiceV2Client =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.discovery.DiscoveryServiceV2PromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.LockRequest,
 *   !proto.discovery.LockResponse>}
 */
const methodDescriptor_DiscoveryServiceV2_Lock = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryServiceV2/Lock',
  grpc.web.MethodType.UNARY,
  proto.discovery.LockRequest,
  proto.discovery.LockResponse,
  /**
   * @param {!proto.discovery.LockRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.LockResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.LockRequest,
 *   !proto.discovery.LockResponse>}
 */
const methodInfo_DiscoveryServiceV2_Lock = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.LockResponse,
  /**
   * @param {!proto.discovery.LockRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.LockResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.LockRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.LockResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.LockResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceV2Client.prototype.lock =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/Lock',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_Lock,
      callback);
};


/**
 * @param {!proto.discovery.LockRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.LockResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServiceV2PromiseClient.prototype.lock =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/Lock',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_Lock);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.RegisterRequest,
 *   !proto.discovery.RegisterResponse>}
 */
const methodDescriptor_DiscoveryServiceV2_RegisterV2 = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryServiceV2/RegisterV2',
  grpc.web.MethodType.UNARY,
  proto.discovery.RegisterRequest,
  proto.discovery.RegisterResponse,
  /**
   * @param {!proto.discovery.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.RegisterResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.RegisterRequest,
 *   !proto.discovery.RegisterResponse>}
 */
const methodInfo_DiscoveryServiceV2_RegisterV2 = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.RegisterResponse,
  /**
   * @param {!proto.discovery.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.RegisterResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.RegisterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.RegisterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceV2Client.prototype.registerV2 =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/RegisterV2',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_RegisterV2,
      callback);
};


/**
 * @param {!proto.discovery.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.RegisterResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServiceV2PromiseClient.prototype.registerV2 =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/RegisterV2',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_RegisterV2);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.GetRequest,
 *   !proto.discovery.GetResponse>}
 */
const methodDescriptor_DiscoveryServiceV2_Get = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryServiceV2/Get',
  grpc.web.MethodType.UNARY,
  proto.discovery.GetRequest,
  proto.discovery.GetResponse,
  /**
   * @param {!proto.discovery.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.GetRequest,
 *   !proto.discovery.GetResponse>}
 */
const methodInfo_DiscoveryServiceV2_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.GetResponse,
  /**
   * @param {!proto.discovery.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.GetResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceV2Client.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/Get',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_Get,
      callback);
};


/**
 * @param {!proto.discovery.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServiceV2PromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/Get',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.UnregisterRequest,
 *   !proto.discovery.UnregisterResponse>}
 */
const methodDescriptor_DiscoveryServiceV2_Unregister = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryServiceV2/Unregister',
  grpc.web.MethodType.UNARY,
  proto.discovery.UnregisterRequest,
  proto.discovery.UnregisterResponse,
  /**
   * @param {!proto.discovery.UnregisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.UnregisterResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.UnregisterRequest,
 *   !proto.discovery.UnregisterResponse>}
 */
const methodInfo_DiscoveryServiceV2_Unregister = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.UnregisterResponse,
  /**
   * @param {!proto.discovery.UnregisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.UnregisterResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.UnregisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.UnregisterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.UnregisterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceV2Client.prototype.unregister =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/Unregister',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_Unregister,
      callback);
};


/**
 * @param {!proto.discovery.UnregisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.UnregisterResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServiceV2PromiseClient.prototype.unregister =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/Unregister',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_Unregister);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.discovery.MasterRequest,
 *   !proto.discovery.MasterResponse>}
 */
const methodDescriptor_DiscoveryServiceV2_MasterElect = new grpc.web.MethodDescriptor(
  '/discovery.DiscoveryServiceV2/MasterElect',
  grpc.web.MethodType.UNARY,
  proto.discovery.MasterRequest,
  proto.discovery.MasterResponse,
  /**
   * @param {!proto.discovery.MasterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.MasterResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.discovery.MasterRequest,
 *   !proto.discovery.MasterResponse>}
 */
const methodInfo_DiscoveryServiceV2_MasterElect = new grpc.web.AbstractClientBase.MethodInfo(
  proto.discovery.MasterResponse,
  /**
   * @param {!proto.discovery.MasterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.discovery.MasterResponse.deserializeBinary
);


/**
 * @param {!proto.discovery.MasterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.discovery.MasterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.discovery.MasterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.discovery.DiscoveryServiceV2Client.prototype.masterElect =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/MasterElect',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_MasterElect,
      callback);
};


/**
 * @param {!proto.discovery.MasterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.discovery.MasterResponse>}
 *     A native promise that resolves to the response
 */
proto.discovery.DiscoveryServiceV2PromiseClient.prototype.masterElect =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/discovery.DiscoveryServiceV2/MasterElect',
      request,
      metadata || {},
      methodDescriptor_DiscoveryServiceV2_MasterElect);
};


module.exports = proto.discovery;

