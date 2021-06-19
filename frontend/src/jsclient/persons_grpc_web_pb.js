/**
 * @fileoverview gRPC-Web generated client stub for backend
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.backend = require('./persons_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.backend.PersonsClient =
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
proto.backend.PersonsPromiseClient =
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
 *   !proto.backend.TestRequest,
 *   !proto.backend.TestResponse>}
 */
const methodDescriptor_Persons_Person = new grpc.web.MethodDescriptor(
  '/backend.Persons/Person',
  grpc.web.MethodType.UNARY,
  proto.backend.TestRequest,
  proto.backend.TestResponse,
  /**
   * @param {!proto.backend.TestRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.backend.TestResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.backend.TestRequest,
 *   !proto.backend.TestResponse>}
 */
const methodInfo_Persons_Person = new grpc.web.AbstractClientBase.MethodInfo(
  proto.backend.TestResponse,
  /**
   * @param {!proto.backend.TestRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.backend.TestResponse.deserializeBinary
);


/**
 * @param {!proto.backend.TestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.backend.TestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.backend.TestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.backend.PersonsClient.prototype.person =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/backend.Persons/Person',
      request,
      metadata || {},
      methodDescriptor_Persons_Person,
      callback);
};


/**
 * @param {!proto.backend.TestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.backend.TestResponse>}
 *     Promise that resolves to the response
 */
proto.backend.PersonsPromiseClient.prototype.person =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/backend.Persons/Person',
      request,
      metadata || {},
      methodDescriptor_Persons_Person);
};


module.exports = proto.backend;

