/*eslint-disable block-scoped-var, no-redeclare, no-control-regex, no-prototype-builtins*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Lazily resolved type references
const $lazyTypes = [];

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.user = (function() {

    /**
     * Namespace user.
     * @exports user
     * @namespace
     */
    const user = {};

    user.UserService = (function() {

        /**
         * Constructs a new UserService service.
         * @exports user.UserService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function UserService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (UserService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = UserService;

        /**
         * Creates new UserService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {UserService} RPC service. Useful where requests and/or responses are streamed.
         */
        UserService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link UserService#createUser}.
         * @typedef UserService_createUser_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {user.CreateUserResponse} [response] CreateUserResponse
         */

        /**
         * Calls CreateUser.
         * @param {user.CreateUserRequest|Object} request CreateUserRequest message or plain object
         * @param {UserService_createUser_Callback} callback Node-style callback called with the error, if any, and CreateUserResponse
         * @returns {undefined}
         */
        UserService.prototype.createUser = function createUser(request, callback) {
            return this.rpcCall(createUser, $root.user.CreateUserRequest, $root.user.CreateUserResponse, request, callback);
        };

        /**
         * Calls CreateUser.
         * @name UserService#createUser
         * @function
         * @param {user.CreateUserRequest|Object} request CreateUserRequest message or plain object
         * @returns {Promise<user.CreateUserResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link UserService#editUser}.
         * @typedef UserService_editUser_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {user.EditUserResponse} [response] EditUserResponse
         */

        /**
         * Calls EditUser.
         * @param {user.EditUserRequest|Object} request EditUserRequest message or plain object
         * @param {UserService_editUser_Callback} callback Node-style callback called with the error, if any, and EditUserResponse
         * @returns {undefined}
         */
        UserService.prototype.editUser = function editUser(request, callback) {
            return this.rpcCall(editUser, $root.user.EditUserRequest, $root.user.EditUserResponse, request, callback);
        };

        /**
         * Calls EditUser.
         * @name UserService#editUser
         * @function
         * @param {user.EditUserRequest|Object} request EditUserRequest message or plain object
         * @returns {Promise<user.EditUserResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link UserService#changeUsername}.
         * @typedef UserService_changeUsername_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {user.ChangeUsernameResponse} [response] ChangeUsernameResponse
         */

        /**
         * Calls ChangeUsername.
         * @param {user.ChangeUsernameRequest|Object} request ChangeUsernameRequest message or plain object
         * @param {UserService_changeUsername_Callback} callback Node-style callback called with the error, if any, and ChangeUsernameResponse
         * @returns {undefined}
         */
        UserService.prototype.changeUsername = function changeUsername(request, callback) {
            return this.rpcCall(changeUsername, $root.user.ChangeUsernameRequest, $root.user.ChangeUsernameResponse, request, callback);
        };

        /**
         * Calls ChangeUsername.
         * @name UserService#changeUsername
         * @function
         * @param {user.ChangeUsernameRequest|Object} request ChangeUsernameRequest message or plain object
         * @returns {Promise<user.ChangeUsernameResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link UserService#deleteUser}.
         * @typedef UserService_deleteUser_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {user.DeleteUserResponse} [response] DeleteUserResponse
         */

        /**
         * Calls DeleteUser.
         * @param {user.DeleteUserRequest|Object} request DeleteUserRequest message or plain object
         * @param {UserService_deleteUser_Callback} callback Node-style callback called with the error, if any, and DeleteUserResponse
         * @returns {undefined}
         */
        UserService.prototype.deleteUser = function deleteUser(request, callback) {
            return this.rpcCall(deleteUser, $root.user.DeleteUserRequest, $root.user.DeleteUserResponse, request, callback);
        };

        /**
         * Calls DeleteUser.
         * @name UserService#deleteUser
         * @function
         * @param {user.DeleteUserRequest|Object} request DeleteUserRequest message or plain object
         * @returns {Promise<user.DeleteUserResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link UserService#resetPassword}.
         * @typedef UserService_resetPassword_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {user.ResetPasswordResponse} [response] ResetPasswordResponse
         */

        /**
         * Calls ResetPassword.
         * @param {user.ResetPasswordRequest|Object} request ResetPasswordRequest message or plain object
         * @param {UserService_resetPassword_Callback} callback Node-style callback called with the error, if any, and ResetPasswordResponse
         * @returns {undefined}
         */
        UserService.prototype.resetPassword = function resetPassword(request, callback) {
            return this.rpcCall(resetPassword, $root.user.ResetPasswordRequest, $root.user.ResetPasswordResponse, request, callback);
        };

        /**
         * Calls ResetPassword.
         * @name UserService#resetPassword
         * @function
         * @param {user.ResetPasswordRequest|Object} request ResetPasswordRequest message or plain object
         * @returns {Promise<user.ResetPasswordResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link UserService#getUser}.
         * @typedef UserService_getUser_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {user.GetUserResponse} [response] GetUserResponse
         */

        /**
         * Calls GetUser.
         * @param {user.GetUserRequest|Object} request GetUserRequest message or plain object
         * @param {UserService_getUser_Callback} callback Node-style callback called with the error, if any, and GetUserResponse
         * @returns {undefined}
         */
        UserService.prototype.getUser = function getUser(request, callback) {
            return this.rpcCall(getUser, $root.user.GetUserRequest, $root.user.GetUserResponse, request, callback);
        };

        /**
         * Calls GetUser.
         * @name UserService#getUser
         * @function
         * @param {user.GetUserRequest|Object} request GetUserRequest message or plain object
         * @returns {Promise<user.GetUserResponse>} Promise
         * @variation 2
         */

        return UserService;
    })();

    user.Address = (function() {

        /**
         * Constructs a new Address.
         * @exports user.Address
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function Address(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * Address ID.
         * @type {number|undefined}
         */
        Address.prototype.ID = 0;

        /**
         * Address postcode.
         * @type {string|undefined}
         */
        Address.prototype.postcode = "";

        /**
         * Address city.
         * @type {string|undefined}
         */
        Address.prototype.city = "";

        /**
         * Address country.
         * @type {string|undefined}
         */
        Address.prototype.country = "";

        /**
         * Address street.
         * @type {string|undefined}
         */
        Address.prototype.street = "";

        /**
         * Address houseno.
         * @type {number|undefined}
         */
        Address.prototype.houseno = 0;

        /**
         * Address additional.
         * @type {string|undefined}
         */
        Address.prototype.additional = "";

        /**
         * Creates a new Address instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.Address} Address instance
         */
        Address.create = function create(properties) {
            return new Address(properties);
        };

        /**
         * Encodes the specified Address message.
         * @param {user.Address|Object} message Address message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Address.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.postcode !== undefined && message.hasOwnProperty("postcode"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.postcode);
            if (message.city !== undefined && message.hasOwnProperty("city"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.city);
            if (message.country !== undefined && message.hasOwnProperty("country"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.country);
            if (message.street !== undefined && message.hasOwnProperty("street"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.street);
            if (message.houseno !== undefined && message.hasOwnProperty("houseno"))
                writer.uint32(/* id 6, wireType 0 =*/48).int32(message.houseno);
            if (message.additional !== undefined && message.hasOwnProperty("additional"))
                writer.uint32(/* id 15, wireType 2 =*/122).string(message.additional);
            return writer;
        };

        /**
         * Encodes the specified Address message, length delimited.
         * @param {user.Address|Object} message Address message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Address.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an Address message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.Address} Address
         */
        Address.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.Address();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.postcode = reader.string();
                    break;
                case 3:
                    message.city = reader.string();
                    break;
                case 4:
                    message.country = reader.string();
                    break;
                case 5:
                    message.street = reader.string();
                    break;
                case 6:
                    message.houseno = reader.int32();
                    break;
                case 15:
                    message.additional = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an Address message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.Address} Address
         */
        Address.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an Address message.
         * @param {user.Address|Object} message Address message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        Address.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.postcode !== undefined)
                if (!$util.isString(message.postcode))
                    return "postcode: string expected";
            if (message.city !== undefined)
                if (!$util.isString(message.city))
                    return "city: string expected";
            if (message.country !== undefined)
                if (!$util.isString(message.country))
                    return "country: string expected";
            if (message.street !== undefined)
                if (!$util.isString(message.street))
                    return "street: string expected";
            if (message.houseno !== undefined)
                if (!$util.isInteger(message.houseno))
                    return "houseno: integer expected";
            if (message.additional !== undefined)
                if (!$util.isString(message.additional))
                    return "additional: string expected";
            return null;
        };

        /**
         * Creates an Address message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Address} Address
         */
        Address.fromObject = function fromObject(object) {
            if (object instanceof $root.user.Address)
                return object;
            let message = new $root.user.Address();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            if (object.postcode !== undefined && object.postcode !== null)
                message.postcode = String(object.postcode);
            if (object.city !== undefined && object.city !== null)
                message.city = String(object.city);
            if (object.country !== undefined && object.country !== null)
                message.country = String(object.country);
            if (object.street !== undefined && object.street !== null)
                message.street = String(object.street);
            if (object.houseno !== undefined && object.houseno !== null)
                message.houseno = object.houseno | 0;
            if (object.additional !== undefined && object.additional !== null)
                message.additional = String(object.additional);
            return message;
        };

        /**
         * Creates an Address message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.Address.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Address} Address
         */
        Address.from = Address.fromObject;

        /**
         * Creates a plain object from an Address message. Also converts values to other types if specified.
         * @param {user.Address} message Address
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Address.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = 0;
                object.postcode = "";
                object.city = "";
                object.country = "";
                object.street = "";
                object.houseno = 0;
                object.additional = "";
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.postcode !== undefined && message.postcode !== null && message.hasOwnProperty("postcode"))
                object.postcode = message.postcode;
            if (message.city !== undefined && message.city !== null && message.hasOwnProperty("city"))
                object.city = message.city;
            if (message.country !== undefined && message.country !== null && message.hasOwnProperty("country"))
                object.country = message.country;
            if (message.street !== undefined && message.street !== null && message.hasOwnProperty("street"))
                object.street = message.street;
            if (message.houseno !== undefined && message.houseno !== null && message.hasOwnProperty("houseno"))
                object.houseno = message.houseno;
            if (message.additional !== undefined && message.additional !== null && message.hasOwnProperty("additional"))
                object.additional = message.additional;
            return object;
        };

        /**
         * Creates a plain object from this Address message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Address.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this Address to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        Address.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Address;
    })();

    user.Config = (function() {

        /**
         * Constructs a new Config.
         * @exports user.Config
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function Config(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * Config admin.
         * @type {boolean|undefined}
         */
        Config.prototype.admin = false;

        /**
         * Config email.
         * @type {string|undefined}
         */
        Config.prototype.email = "";

        /**
         * Config password.
         * @type {string|undefined}
         */
        Config.prototype.password = "";

        /**
         * Config salt.
         * @type {string|undefined}
         */
        Config.prototype.salt = "";

        /**
         * Config Address.
         * @type {user.Address|undefined}
         */
        Config.prototype.Address = null;

        /**
         * Config addressID.
         * @type {number|undefined}
         */
        Config.prototype.addressID = 0;

        /**
         * Config phone.
         * @type {string|undefined}
         */
        Config.prototype.phone = "";

        /**
         * Config image.
         * @type {string|undefined}
         */
        Config.prototype.image = "";

        // Lazily resolved type references
        const $types = {
            4: "user.Address"
        }; $lazyTypes.push($types);

        /**
         * Creates a new Config instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.Config} Config instance
         */
        Config.create = function create(properties) {
            return new Config(properties);
        };

        /**
         * Encodes the specified Config message.
         * @param {user.Config|Object} message Config message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Config.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.admin !== undefined && message.hasOwnProperty("admin"))
                writer.uint32(/* id 1, wireType 0 =*/8).bool(message.admin);
            if (message.email !== undefined && message.hasOwnProperty("email"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.email);
            if (message.password !== undefined && message.hasOwnProperty("password"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.password);
            if (message.salt !== undefined && message.hasOwnProperty("salt"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.salt);
            if (message.Address && message.hasOwnProperty("Address"))
                $types[4].encode(message.Address, writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            if (message.addressID !== undefined && message.hasOwnProperty("addressID"))
                writer.uint32(/* id 6, wireType 0 =*/48).uint32(message.addressID);
            if (message.phone !== undefined && message.hasOwnProperty("phone"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.phone);
            if (message.image !== undefined && message.hasOwnProperty("image"))
                writer.uint32(/* id 15, wireType 2 =*/122).string(message.image);
            return writer;
        };

        /**
         * Encodes the specified Config message, length delimited.
         * @param {user.Config|Object} message Config message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Config.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Config message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.Config} Config
         */
        Config.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.Config();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.admin = reader.bool();
                    break;
                case 2:
                    message.email = reader.string();
                    break;
                case 3:
                    message.password = reader.string();
                    break;
                case 4:
                    message.salt = reader.string();
                    break;
                case 5:
                    message.Address = $types[4].decode(reader, reader.uint32());
                    break;
                case 6:
                    message.addressID = reader.uint32();
                    break;
                case 7:
                    message.phone = reader.string();
                    break;
                case 15:
                    message.image = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Config message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.Config} Config
         */
        Config.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Config message.
         * @param {user.Config|Object} message Config message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        Config.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.admin !== undefined)
                if (typeof message.admin !== "boolean")
                    return "admin: boolean expected";
            if (message.email !== undefined)
                if (!$util.isString(message.email))
                    return "email: string expected";
            if (message.password !== undefined)
                if (!$util.isString(message.password))
                    return "password: string expected";
            if (message.salt !== undefined)
                if (!$util.isString(message.salt))
                    return "salt: string expected";
            if (message.Address !== undefined && message.Address !== null) {
                let error = $types[4].verify(message.Address);
                if (error)
                    return "Address." + error;
            }
            if (message.addressID !== undefined)
                if (!$util.isInteger(message.addressID))
                    return "addressID: integer expected";
            if (message.phone !== undefined)
                if (!$util.isString(message.phone))
                    return "phone: string expected";
            if (message.image !== undefined)
                if (!$util.isString(message.image))
                    return "image: string expected";
            return null;
        };

        /**
         * Creates a Config message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Config} Config
         */
        Config.fromObject = function fromObject(object) {
            if (object instanceof $root.user.Config)
                return object;
            let message = new $root.user.Config();
            if (object.admin !== undefined && object.admin !== null)
                message.admin = Boolean(object.admin);
            if (object.email !== undefined && object.email !== null)
                message.email = String(object.email);
            if (object.password !== undefined && object.password !== null)
                message.password = String(object.password);
            if (object.salt !== undefined && object.salt !== null)
                message.salt = String(object.salt);
            if (object.Address !== undefined && object.Address !== null) {
                if (typeof object.Address !== "object")
                    throw TypeError(".user.Config.Address: object expected");
                message.Address = $types[4].fromObject(object.Address);
            }
            if (object.addressID !== undefined && object.addressID !== null)
                message.addressID = object.addressID >>> 0;
            if (object.phone !== undefined && object.phone !== null)
                message.phone = String(object.phone);
            if (object.image !== undefined && object.image !== null)
                message.image = String(object.image);
            return message;
        };

        /**
         * Creates a Config message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.Config.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Config} Config
         */
        Config.from = Config.fromObject;

        /**
         * Creates a plain object from a Config message. Also converts values to other types if specified.
         * @param {user.Config} message Config
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Config.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.admin = false;
                object.email = "";
                object.password = "";
                object.salt = "";
                object.Address = null;
                object.addressID = 0;
                object.phone = "";
                object.image = "";
            }
            if (message.admin !== undefined && message.admin !== null && message.hasOwnProperty("admin"))
                object.admin = message.admin;
            if (message.email !== undefined && message.email !== null && message.hasOwnProperty("email"))
                object.email = message.email;
            if (message.password !== undefined && message.password !== null && message.hasOwnProperty("password"))
                object.password = message.password;
            if (message.salt !== undefined && message.salt !== null && message.hasOwnProperty("salt"))
                object.salt = message.salt;
            if (message.Address !== undefined && message.Address !== null && message.hasOwnProperty("Address"))
                object.Address = $types[4].toObject(message.Address, options);
            if (message.addressID !== undefined && message.addressID !== null && message.hasOwnProperty("addressID"))
                object.addressID = message.addressID;
            if (message.phone !== undefined && message.phone !== null && message.hasOwnProperty("phone"))
                object.phone = message.phone;
            if (message.image !== undefined && message.image !== null && message.hasOwnProperty("image"))
                object.image = message.image;
            return object;
        };

        /**
         * Creates a plain object from this Config message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Config.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this Config to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        Config.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Config;
    })();

    user.User = (function() {

        /**
         * Constructs a new User.
         * @exports user.User
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function User(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * User ID.
         * @type {number|undefined}
         */
        User.prototype.ID = 0;

        /**
         * User username.
         * @type {string|undefined}
         */
        User.prototype.username = "";

        /**
         * User config.
         * @type {user.Config|undefined}
         */
        User.prototype.config = null;

        // Lazily resolved type references
        const $types = {
            2: "user.Config"
        }; $lazyTypes.push($types);

        /**
         * Creates a new User instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.User} User instance
         */
        User.create = function create(properties) {
            return new User(properties);
        };

        /**
         * Encodes the specified User message.
         * @param {user.User|Object} message User message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        User.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.username !== undefined && message.hasOwnProperty("username"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.username);
            if (message.config && message.hasOwnProperty("config"))
                $types[2].encode(message.config, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified User message, length delimited.
         * @param {user.User|Object} message User message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        User.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a User message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.User} User
         */
        User.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.User();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.username = reader.string();
                    break;
                case 3:
                    message.config = $types[2].decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a User message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.User} User
         */
        User.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a User message.
         * @param {user.User|Object} message User message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        User.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.username !== undefined)
                if (!$util.isString(message.username))
                    return "username: string expected";
            if (message.config !== undefined && message.config !== null) {
                let error = $types[2].verify(message.config);
                if (error)
                    return "config." + error;
            }
            return null;
        };

        /**
         * Creates a User message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.User} User
         */
        User.fromObject = function fromObject(object) {
            if (object instanceof $root.user.User)
                return object;
            let message = new $root.user.User();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            if (object.username !== undefined && object.username !== null)
                message.username = String(object.username);
            if (object.config !== undefined && object.config !== null) {
                if (typeof object.config !== "object")
                    throw TypeError(".user.User.config: object expected");
                message.config = $types[2].fromObject(object.config);
            }
            return message;
        };

        /**
         * Creates a User message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.User.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.User} User
         */
        User.from = User.fromObject;

        /**
         * Creates a plain object from a User message. Also converts values to other types if specified.
         * @param {user.User} message User
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        User.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = 0;
                object.username = "";
                object.config = null;
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.username !== undefined && message.username !== null && message.hasOwnProperty("username"))
                object.username = message.username;
            if (message.config !== undefined && message.config !== null && message.hasOwnProperty("config"))
                object.config = $types[2].toObject(message.config, options);
            return object;
        };

        /**
         * Creates a plain object from this User message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        User.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this User to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        User.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return User;
    })();

    user.CreateUserRequest = (function() {

        /**
         * Constructs a new CreateUserRequest.
         * @exports user.CreateUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function CreateUserRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * CreateUserRequest username.
         * @type {string|undefined}
         */
        CreateUserRequest.prototype.username = "";

        /**
         * CreateUserRequest config.
         * @type {user.Config|undefined}
         */
        CreateUserRequest.prototype.config = null;

        // Lazily resolved type references
        const $types = {
            1: "user.Config"
        }; $lazyTypes.push($types);

        /**
         * Creates a new CreateUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.CreateUserRequest} CreateUserRequest instance
         */
        CreateUserRequest.create = function create(properties) {
            return new CreateUserRequest(properties);
        };

        /**
         * Encodes the specified CreateUserRequest message.
         * @param {user.CreateUserRequest|Object} message CreateUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateUserRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.username !== undefined && message.hasOwnProperty("username"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.username);
            if (message.config && message.hasOwnProperty("config"))
                $types[1].encode(message.config, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified CreateUserRequest message, length delimited.
         * @param {user.CreateUserRequest|Object} message CreateUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateUserRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CreateUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        CreateUserRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.CreateUserRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.username = reader.string();
                    break;
                case 2:
                    message.config = $types[1].decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CreateUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        CreateUserRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CreateUserRequest message.
         * @param {user.CreateUserRequest|Object} message CreateUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        CreateUserRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.username !== undefined)
                if (!$util.isString(message.username))
                    return "username: string expected";
            if (message.config !== undefined && message.config !== null) {
                let error = $types[1].verify(message.config);
                if (error)
                    return "config." + error;
            }
            return null;
        };

        /**
         * Creates a CreateUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        CreateUserRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.user.CreateUserRequest)
                return object;
            let message = new $root.user.CreateUserRequest();
            if (object.username !== undefined && object.username !== null)
                message.username = String(object.username);
            if (object.config !== undefined && object.config !== null) {
                if (typeof object.config !== "object")
                    throw TypeError(".user.CreateUserRequest.config: object expected");
                message.config = $types[1].fromObject(object.config);
            }
            return message;
        };

        /**
         * Creates a CreateUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.CreateUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        CreateUserRequest.from = CreateUserRequest.fromObject;

        /**
         * Creates a plain object from a CreateUserRequest message. Also converts values to other types if specified.
         * @param {user.CreateUserRequest} message CreateUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateUserRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.username = "";
                object.config = null;
            }
            if (message.username !== undefined && message.username !== null && message.hasOwnProperty("username"))
                object.username = message.username;
            if (message.config !== undefined && message.config !== null && message.hasOwnProperty("config"))
                object.config = $types[1].toObject(message.config, options);
            return object;
        };

        /**
         * Creates a plain object from this CreateUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateUserRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this CreateUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        CreateUserRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CreateUserRequest;
    })();

    user.CreateUserResponse = (function() {

        /**
         * Constructs a new CreateUserResponse.
         * @exports user.CreateUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function CreateUserResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * CreateUserResponse ID.
         * @type {number|undefined}
         */
        CreateUserResponse.prototype.ID = 0;

        /**
         * CreateUserResponse error.
         * @type {string|undefined}
         */
        CreateUserResponse.prototype.error = "";

        /**
         * Creates a new CreateUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.CreateUserResponse} CreateUserResponse instance
         */
        CreateUserResponse.create = function create(properties) {
            return new CreateUserResponse(properties);
        };

        /**
         * Encodes the specified CreateUserResponse message.
         * @param {user.CreateUserResponse|Object} message CreateUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateUserResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified CreateUserResponse message, length delimited.
         * @param {user.CreateUserResponse|Object} message CreateUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateUserResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CreateUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        CreateUserResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.CreateUserResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CreateUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        CreateUserResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CreateUserResponse message.
         * @param {user.CreateUserResponse|Object} message CreateUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        CreateUserResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a CreateUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        CreateUserResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.user.CreateUserResponse)
                return object;
            let message = new $root.user.CreateUserResponse();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a CreateUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.CreateUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        CreateUserResponse.from = CreateUserResponse.fromObject;

        /**
         * Creates a plain object from a CreateUserResponse message. Also converts values to other types if specified.
         * @param {user.CreateUserResponse} message CreateUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateUserResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = 0;
                object.error = "";
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this CreateUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateUserResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this CreateUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        CreateUserResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CreateUserResponse;
    })();

    user.EditUserRequest = (function() {

        /**
         * Constructs a new EditUserRequest.
         * @exports user.EditUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function EditUserRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * EditUserRequest ID.
         * @type {number|undefined}
         */
        EditUserRequest.prototype.ID = 0;

        /**
         * EditUserRequest config.
         * @type {user.Config|undefined}
         */
        EditUserRequest.prototype.config = null;

        // Lazily resolved type references
        const $types = {
            1: "user.Config"
        }; $lazyTypes.push($types);

        /**
         * Creates a new EditUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.EditUserRequest} EditUserRequest instance
         */
        EditUserRequest.create = function create(properties) {
            return new EditUserRequest(properties);
        };

        /**
         * Encodes the specified EditUserRequest message.
         * @param {user.EditUserRequest|Object} message EditUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditUserRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.config && message.hasOwnProperty("config"))
                $types[1].encode(message.config, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified EditUserRequest message, length delimited.
         * @param {user.EditUserRequest|Object} message EditUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditUserRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an EditUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.EditUserRequest} EditUserRequest
         */
        EditUserRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.EditUserRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.config = $types[1].decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an EditUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.EditUserRequest} EditUserRequest
         */
        EditUserRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an EditUserRequest message.
         * @param {user.EditUserRequest|Object} message EditUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        EditUserRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.config !== undefined && message.config !== null) {
                let error = $types[1].verify(message.config);
                if (error)
                    return "config." + error;
            }
            return null;
        };

        /**
         * Creates an EditUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserRequest} EditUserRequest
         */
        EditUserRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.user.EditUserRequest)
                return object;
            let message = new $root.user.EditUserRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            if (object.config !== undefined && object.config !== null) {
                if (typeof object.config !== "object")
                    throw TypeError(".user.EditUserRequest.config: object expected");
                message.config = $types[1].fromObject(object.config);
            }
            return message;
        };

        /**
         * Creates an EditUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.EditUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserRequest} EditUserRequest
         */
        EditUserRequest.from = EditUserRequest.fromObject;

        /**
         * Creates a plain object from an EditUserRequest message. Also converts values to other types if specified.
         * @param {user.EditUserRequest} message EditUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditUserRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = 0;
                object.config = null;
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.config !== undefined && message.config !== null && message.hasOwnProperty("config"))
                object.config = $types[1].toObject(message.config, options);
            return object;
        };

        /**
         * Creates a plain object from this EditUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditUserRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this EditUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        EditUserRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return EditUserRequest;
    })();

    user.EditUserResponse = (function() {

        /**
         * Constructs a new EditUserResponse.
         * @exports user.EditUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function EditUserResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * EditUserResponse error.
         * @type {string|undefined}
         */
        EditUserResponse.prototype.error = "";

        /**
         * Creates a new EditUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.EditUserResponse} EditUserResponse instance
         */
        EditUserResponse.create = function create(properties) {
            return new EditUserResponse(properties);
        };

        /**
         * Encodes the specified EditUserResponse message.
         * @param {user.EditUserResponse|Object} message EditUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditUserResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified EditUserResponse message, length delimited.
         * @param {user.EditUserResponse|Object} message EditUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditUserResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an EditUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.EditUserResponse} EditUserResponse
         */
        EditUserResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.EditUserResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an EditUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.EditUserResponse} EditUserResponse
         */
        EditUserResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an EditUserResponse message.
         * @param {user.EditUserResponse|Object} message EditUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        EditUserResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates an EditUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserResponse} EditUserResponse
         */
        EditUserResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.user.EditUserResponse)
                return object;
            let message = new $root.user.EditUserResponse();
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates an EditUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.EditUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserResponse} EditUserResponse
         */
        EditUserResponse.from = EditUserResponse.fromObject;

        /**
         * Creates a plain object from an EditUserResponse message. Also converts values to other types if specified.
         * @param {user.EditUserResponse} message EditUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditUserResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.error = "";
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this EditUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditUserResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this EditUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        EditUserResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return EditUserResponse;
    })();

    user.ChangeUsernameRequest = (function() {

        /**
         * Constructs a new ChangeUsernameRequest.
         * @exports user.ChangeUsernameRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function ChangeUsernameRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * ChangeUsernameRequest ID.
         * @type {number|undefined}
         */
        ChangeUsernameRequest.prototype.ID = 0;

        /**
         * ChangeUsernameRequest username.
         * @type {string|undefined}
         */
        ChangeUsernameRequest.prototype.username = "";

        /**
         * Creates a new ChangeUsernameRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest instance
         */
        ChangeUsernameRequest.create = function create(properties) {
            return new ChangeUsernameRequest(properties);
        };

        /**
         * Encodes the specified ChangeUsernameRequest message.
         * @param {user.ChangeUsernameRequest|Object} message ChangeUsernameRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ChangeUsernameRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.username !== undefined && message.hasOwnProperty("username"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.username);
            return writer;
        };

        /**
         * Encodes the specified ChangeUsernameRequest message, length delimited.
         * @param {user.ChangeUsernameRequest|Object} message ChangeUsernameRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ChangeUsernameRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ChangeUsernameRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        ChangeUsernameRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.ChangeUsernameRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.username = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ChangeUsernameRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        ChangeUsernameRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ChangeUsernameRequest message.
         * @param {user.ChangeUsernameRequest|Object} message ChangeUsernameRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        ChangeUsernameRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.username !== undefined)
                if (!$util.isString(message.username))
                    return "username: string expected";
            return null;
        };

        /**
         * Creates a ChangeUsernameRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        ChangeUsernameRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.user.ChangeUsernameRequest)
                return object;
            let message = new $root.user.ChangeUsernameRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            if (object.username !== undefined && object.username !== null)
                message.username = String(object.username);
            return message;
        };

        /**
         * Creates a ChangeUsernameRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ChangeUsernameRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        ChangeUsernameRequest.from = ChangeUsernameRequest.fromObject;

        /**
         * Creates a plain object from a ChangeUsernameRequest message. Also converts values to other types if specified.
         * @param {user.ChangeUsernameRequest} message ChangeUsernameRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ChangeUsernameRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = 0;
                object.username = "";
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.username !== undefined && message.username !== null && message.hasOwnProperty("username"))
                object.username = message.username;
            return object;
        };

        /**
         * Creates a plain object from this ChangeUsernameRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ChangeUsernameRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this ChangeUsernameRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        ChangeUsernameRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ChangeUsernameRequest;
    })();

    user.ChangeUsernameResponse = (function() {

        /**
         * Constructs a new ChangeUsernameResponse.
         * @exports user.ChangeUsernameResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function ChangeUsernameResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * ChangeUsernameResponse error.
         * @type {string|undefined}
         */
        ChangeUsernameResponse.prototype.error = "";

        /**
         * Creates a new ChangeUsernameResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse instance
         */
        ChangeUsernameResponse.create = function create(properties) {
            return new ChangeUsernameResponse(properties);
        };

        /**
         * Encodes the specified ChangeUsernameResponse message.
         * @param {user.ChangeUsernameResponse|Object} message ChangeUsernameResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ChangeUsernameResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified ChangeUsernameResponse message, length delimited.
         * @param {user.ChangeUsernameResponse|Object} message ChangeUsernameResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ChangeUsernameResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ChangeUsernameResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        ChangeUsernameResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.ChangeUsernameResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ChangeUsernameResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        ChangeUsernameResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ChangeUsernameResponse message.
         * @param {user.ChangeUsernameResponse|Object} message ChangeUsernameResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        ChangeUsernameResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a ChangeUsernameResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        ChangeUsernameResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.user.ChangeUsernameResponse)
                return object;
            let message = new $root.user.ChangeUsernameResponse();
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a ChangeUsernameResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ChangeUsernameResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        ChangeUsernameResponse.from = ChangeUsernameResponse.fromObject;

        /**
         * Creates a plain object from a ChangeUsernameResponse message. Also converts values to other types if specified.
         * @param {user.ChangeUsernameResponse} message ChangeUsernameResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ChangeUsernameResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.error = "";
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this ChangeUsernameResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ChangeUsernameResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this ChangeUsernameResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        ChangeUsernameResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ChangeUsernameResponse;
    })();

    user.DeleteUserRequest = (function() {

        /**
         * Constructs a new DeleteUserRequest.
         * @exports user.DeleteUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function DeleteUserRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * DeleteUserRequest ID.
         * @type {number|undefined}
         */
        DeleteUserRequest.prototype.ID = 0;

        /**
         * Creates a new DeleteUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.DeleteUserRequest} DeleteUserRequest instance
         */
        DeleteUserRequest.create = function create(properties) {
            return new DeleteUserRequest(properties);
        };

        /**
         * Encodes the specified DeleteUserRequest message.
         * @param {user.DeleteUserRequest|Object} message DeleteUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DeleteUserRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            return writer;
        };

        /**
         * Encodes the specified DeleteUserRequest message, length delimited.
         * @param {user.DeleteUserRequest|Object} message DeleteUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DeleteUserRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DeleteUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        DeleteUserRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.DeleteUserRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DeleteUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        DeleteUserRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DeleteUserRequest message.
         * @param {user.DeleteUserRequest|Object} message DeleteUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        DeleteUserRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            return null;
        };

        /**
         * Creates a DeleteUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        DeleteUserRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.user.DeleteUserRequest)
                return object;
            let message = new $root.user.DeleteUserRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            return message;
        };

        /**
         * Creates a DeleteUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.DeleteUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        DeleteUserRequest.from = DeleteUserRequest.fromObject;

        /**
         * Creates a plain object from a DeleteUserRequest message. Also converts values to other types if specified.
         * @param {user.DeleteUserRequest} message DeleteUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DeleteUserRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ID = 0;
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            return object;
        };

        /**
         * Creates a plain object from this DeleteUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DeleteUserRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this DeleteUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        DeleteUserRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DeleteUserRequest;
    })();

    user.DeleteUserResponse = (function() {

        /**
         * Constructs a new DeleteUserResponse.
         * @exports user.DeleteUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function DeleteUserResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * DeleteUserResponse error.
         * @type {string|undefined}
         */
        DeleteUserResponse.prototype.error = "";

        /**
         * Creates a new DeleteUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.DeleteUserResponse} DeleteUserResponse instance
         */
        DeleteUserResponse.create = function create(properties) {
            return new DeleteUserResponse(properties);
        };

        /**
         * Encodes the specified DeleteUserResponse message.
         * @param {user.DeleteUserResponse|Object} message DeleteUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DeleteUserResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified DeleteUserResponse message, length delimited.
         * @param {user.DeleteUserResponse|Object} message DeleteUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DeleteUserResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DeleteUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        DeleteUserResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.DeleteUserResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DeleteUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        DeleteUserResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DeleteUserResponse message.
         * @param {user.DeleteUserResponse|Object} message DeleteUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        DeleteUserResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a DeleteUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        DeleteUserResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.user.DeleteUserResponse)
                return object;
            let message = new $root.user.DeleteUserResponse();
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a DeleteUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.DeleteUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        DeleteUserResponse.from = DeleteUserResponse.fromObject;

        /**
         * Creates a plain object from a DeleteUserResponse message. Also converts values to other types if specified.
         * @param {user.DeleteUserResponse} message DeleteUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DeleteUserResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.error = "";
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this DeleteUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DeleteUserResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this DeleteUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        DeleteUserResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DeleteUserResponse;
    })();

    user.ResetPasswordRequest = (function() {

        /**
         * Constructs a new ResetPasswordRequest.
         * @exports user.ResetPasswordRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function ResetPasswordRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * ResetPasswordRequest email.
         * @type {string|undefined}
         */
        ResetPasswordRequest.prototype.email = "";

        /**
         * Creates a new ResetPasswordRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest instance
         */
        ResetPasswordRequest.create = function create(properties) {
            return new ResetPasswordRequest(properties);
        };

        /**
         * Encodes the specified ResetPasswordRequest message.
         * @param {user.ResetPasswordRequest|Object} message ResetPasswordRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ResetPasswordRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.email !== undefined && message.hasOwnProperty("email"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.email);
            return writer;
        };

        /**
         * Encodes the specified ResetPasswordRequest message, length delimited.
         * @param {user.ResetPasswordRequest|Object} message ResetPasswordRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ResetPasswordRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ResetPasswordRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        ResetPasswordRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.ResetPasswordRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.email = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ResetPasswordRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        ResetPasswordRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ResetPasswordRequest message.
         * @param {user.ResetPasswordRequest|Object} message ResetPasswordRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        ResetPasswordRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.email !== undefined)
                if (!$util.isString(message.email))
                    return "email: string expected";
            return null;
        };

        /**
         * Creates a ResetPasswordRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        ResetPasswordRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.user.ResetPasswordRequest)
                return object;
            let message = new $root.user.ResetPasswordRequest();
            if (object.email !== undefined && object.email !== null)
                message.email = String(object.email);
            return message;
        };

        /**
         * Creates a ResetPasswordRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ResetPasswordRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        ResetPasswordRequest.from = ResetPasswordRequest.fromObject;

        /**
         * Creates a plain object from a ResetPasswordRequest message. Also converts values to other types if specified.
         * @param {user.ResetPasswordRequest} message ResetPasswordRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ResetPasswordRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.email = "";
            if (message.email !== undefined && message.email !== null && message.hasOwnProperty("email"))
                object.email = message.email;
            return object;
        };

        /**
         * Creates a plain object from this ResetPasswordRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ResetPasswordRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this ResetPasswordRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        ResetPasswordRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ResetPasswordRequest;
    })();

    user.ResetPasswordResponse = (function() {

        /**
         * Constructs a new ResetPasswordResponse.
         * @exports user.ResetPasswordResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function ResetPasswordResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * ResetPasswordResponse error.
         * @type {string|undefined}
         */
        ResetPasswordResponse.prototype.error = "";

        /**
         * Creates a new ResetPasswordResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse instance
         */
        ResetPasswordResponse.create = function create(properties) {
            return new ResetPasswordResponse(properties);
        };

        /**
         * Encodes the specified ResetPasswordResponse message.
         * @param {user.ResetPasswordResponse|Object} message ResetPasswordResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ResetPasswordResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified ResetPasswordResponse message, length delimited.
         * @param {user.ResetPasswordResponse|Object} message ResetPasswordResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ResetPasswordResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ResetPasswordResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        ResetPasswordResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.ResetPasswordResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ResetPasswordResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        ResetPasswordResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ResetPasswordResponse message.
         * @param {user.ResetPasswordResponse|Object} message ResetPasswordResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        ResetPasswordResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a ResetPasswordResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        ResetPasswordResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.user.ResetPasswordResponse)
                return object;
            let message = new $root.user.ResetPasswordResponse();
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a ResetPasswordResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ResetPasswordResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        ResetPasswordResponse.from = ResetPasswordResponse.fromObject;

        /**
         * Creates a plain object from a ResetPasswordResponse message. Also converts values to other types if specified.
         * @param {user.ResetPasswordResponse} message ResetPasswordResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ResetPasswordResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.error = "";
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this ResetPasswordResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ResetPasswordResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this ResetPasswordResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        ResetPasswordResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ResetPasswordResponse;
    })();

    user.GetUserRequest = (function() {

        /**
         * Constructs a new GetUserRequest.
         * @exports user.GetUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function GetUserRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetUserRequest ID.
         * @type {number|undefined}
         */
        GetUserRequest.prototype.ID = 0;

        /**
         * Creates a new GetUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.GetUserRequest} GetUserRequest instance
         */
        GetUserRequest.create = function create(properties) {
            return new GetUserRequest(properties);
        };

        /**
         * Encodes the specified GetUserRequest message.
         * @param {user.GetUserRequest|Object} message GetUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetUserRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            return writer;
        };

        /**
         * Encodes the specified GetUserRequest message, length delimited.
         * @param {user.GetUserRequest|Object} message GetUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetUserRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GetUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.GetUserRequest} GetUserRequest
         */
        GetUserRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.GetUserRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GetUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.GetUserRequest} GetUserRequest
         */
        GetUserRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GetUserRequest message.
         * @param {user.GetUserRequest|Object} message GetUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        GetUserRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            return null;
        };

        /**
         * Creates a GetUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserRequest} GetUserRequest
         */
        GetUserRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.user.GetUserRequest)
                return object;
            let message = new $root.user.GetUserRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            return message;
        };

        /**
         * Creates a GetUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.GetUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserRequest} GetUserRequest
         */
        GetUserRequest.from = GetUserRequest.fromObject;

        /**
         * Creates a plain object from a GetUserRequest message. Also converts values to other types if specified.
         * @param {user.GetUserRequest} message GetUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetUserRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ID = 0;
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            return object;
        };

        /**
         * Creates a plain object from this GetUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetUserRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this GetUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        GetUserRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return GetUserRequest;
    })();

    user.GetUserResponse = (function() {

        /**
         * Constructs a new GetUserResponse.
         * @exports user.GetUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function GetUserResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetUserResponse user.
         * @type {user.User|undefined}
         */
        GetUserResponse.prototype.user = null;

        /**
         * GetUserResponse error.
         * @type {string|undefined}
         */
        GetUserResponse.prototype.error = "";

        // Lazily resolved type references
        const $types = {
            0: "user.User"
        }; $lazyTypes.push($types);

        /**
         * Creates a new GetUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.GetUserResponse} GetUserResponse instance
         */
        GetUserResponse.create = function create(properties) {
            return new GetUserResponse(properties);
        };

        /**
         * Encodes the specified GetUserResponse message.
         * @param {user.GetUserResponse|Object} message GetUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetUserResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.user && message.hasOwnProperty("user"))
                $types[0].encode(message.user, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified GetUserResponse message, length delimited.
         * @param {user.GetUserResponse|Object} message GetUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetUserResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GetUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.GetUserResponse} GetUserResponse
         */
        GetUserResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.user.GetUserResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.user = $types[0].decode(reader, reader.uint32());
                    break;
                case 2:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GetUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.GetUserResponse} GetUserResponse
         */
        GetUserResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GetUserResponse message.
         * @param {user.GetUserResponse|Object} message GetUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        GetUserResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.user !== undefined && message.user !== null) {
                let error = $types[0].verify(message.user);
                if (error)
                    return "user." + error;
            }
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a GetUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserResponse} GetUserResponse
         */
        GetUserResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.user.GetUserResponse)
                return object;
            let message = new $root.user.GetUserResponse();
            if (object.user !== undefined && object.user !== null) {
                if (typeof object.user !== "object")
                    throw TypeError(".user.GetUserResponse.user: object expected");
                message.user = $types[0].fromObject(object.user);
            }
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a GetUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.GetUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserResponse} GetUserResponse
         */
        GetUserResponse.from = GetUserResponse.fromObject;

        /**
         * Creates a plain object from a GetUserResponse message. Also converts values to other types if specified.
         * @param {user.GetUserResponse} message GetUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetUserResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.user = null;
                object.error = "";
            }
            if (message.user !== undefined && message.user !== null && message.hasOwnProperty("user"))
                object.user = $types[0].toObject(message.user, options);
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this GetUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetUserResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this GetUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        GetUserResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return GetUserResponse;
    })();

    return user;
})();

$root.kmi = (function() {

    /**
     * Namespace kmi.
     * @exports kmi
     * @namespace
     */
    const kmi = {};

    kmi.KMIService = (function() {

        /**
         * Constructs a new KMIService service.
         * @exports kmi.KMIService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function KMIService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (KMIService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = KMIService;

        /**
         * Creates new KMIService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {KMIService} RPC service. Useful where requests and/or responses are streamed.
         */
        KMIService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link KMIService#addKMI}.
         * @typedef KMIService_addKMI_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {kmi.AddKMIResponse} [response] AddKMIResponse
         */

        /**
         * Calls AddKMI.
         * @param {kmi.AddKMIRequest|Object} request AddKMIRequest message or plain object
         * @param {KMIService_addKMI_Callback} callback Node-style callback called with the error, if any, and AddKMIResponse
         * @returns {undefined}
         */
        KMIService.prototype.addKMI = function addKMI(request, callback) {
            return this.rpcCall(addKMI, $root.kmi.AddKMIRequest, $root.kmi.AddKMIResponse, request, callback);
        };

        /**
         * Calls AddKMI.
         * @name KMIService#addKMI
         * @function
         * @param {kmi.AddKMIRequest|Object} request AddKMIRequest message or plain object
         * @returns {Promise<kmi.AddKMIResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link KMIService#removeKMI}.
         * @typedef KMIService_removeKMI_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {kmi.RemoveKMIResponse} [response] RemoveKMIResponse
         */

        /**
         * Calls RemoveKMI.
         * @param {kmi.RemoveKMIRequest|Object} request RemoveKMIRequest message or plain object
         * @param {KMIService_removeKMI_Callback} callback Node-style callback called with the error, if any, and RemoveKMIResponse
         * @returns {undefined}
         */
        KMIService.prototype.removeKMI = function removeKMI(request, callback) {
            return this.rpcCall(removeKMI, $root.kmi.RemoveKMIRequest, $root.kmi.RemoveKMIResponse, request, callback);
        };

        /**
         * Calls RemoveKMI.
         * @name KMIService#removeKMI
         * @function
         * @param {kmi.RemoveKMIRequest|Object} request RemoveKMIRequest message or plain object
         * @returns {Promise<kmi.RemoveKMIResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link KMIService#getKMI}.
         * @typedef KMIService_getKMI_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {kmi.GetKMIResponse} [response] GetKMIResponse
         */

        /**
         * Calls GetKMI.
         * @param {kmi.GetKMIRequest|Object} request GetKMIRequest message or plain object
         * @param {KMIService_getKMI_Callback} callback Node-style callback called with the error, if any, and GetKMIResponse
         * @returns {undefined}
         */
        KMIService.prototype.getKMI = function getKMI(request, callback) {
            return this.rpcCall(getKMI, $root.kmi.GetKMIRequest, $root.kmi.GetKMIResponse, request, callback);
        };

        /**
         * Calls GetKMI.
         * @name KMIService#getKMI
         * @function
         * @param {kmi.GetKMIRequest|Object} request GetKMIRequest message or plain object
         * @returns {Promise<kmi.GetKMIResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link KMIService#kMI}.
         * @typedef KMIService_kMI_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {kmi.KMIResponse} [response] KMIResponse
         */

        /**
         * Calls KMI.
         * @param {kmi.KMIRequest|Object} request KMIRequest message or plain object
         * @param {KMIService_kMI_Callback} callback Node-style callback called with the error, if any, and KMIResponse
         * @returns {undefined}
         */
        KMIService.prototype.kMI = function kMI(request, callback) {
            return this.rpcCall(kMI, $root.kmi.KMIRequest, $root.kmi.KMIResponse, request, callback);
        };

        /**
         * Calls KMI.
         * @name KMIService#kMI
         * @function
         * @param {kmi.KMIRequest|Object} request KMIRequest message or plain object
         * @returns {Promise<kmi.KMIResponse>} Promise
         * @variation 2
         */

        return KMIService;
    })();

    /**
     * Type enum.
     * @name Type
     * @memberof kmi
     * @enum {number}
     * @property {number} WEBSERVER=0 WEBSERVER value
     */
    kmi.Type = (function() {
        const valuesById = {}, values = Object.create(valuesById);
        values["WEBSERVER"] = 0;
        return values;
    })();

    kmi.KMDI = (function() {

        /**
         * Constructs a new KMDI.
         * @exports kmi.KMDI
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function KMDI(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * KMDI ID.
         * @type {number|undefined}
         */
        KMDI.prototype.ID = 0;

        /**
         * KMDI name.
         * @type {string|undefined}
         */
        KMDI.prototype.name = "";

        /**
         * KMDI version.
         * @type {string|undefined}
         */
        KMDI.prototype.version = "";

        /**
         * KMDI description.
         * @type {string|undefined}
         */
        KMDI.prototype.description = "";

        /**
         * KMDI type.
         * @type {number|undefined}
         */
        KMDI.prototype.type = 0;

        // Lazily resolved type references
        const $types = {
            4: "kmi.Type"
        }; $lazyTypes.push($types);

        /**
         * Creates a new KMDI instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMDI} KMDI instance
         */
        KMDI.create = function create(properties) {
            return new KMDI(properties);
        };

        /**
         * Encodes the specified KMDI message.
         * @param {kmi.KMDI|Object} message KMDI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMDI.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.name !== undefined && message.hasOwnProperty("name"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
            if (message.version !== undefined && message.hasOwnProperty("version"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.version);
            if (message.description !== undefined && message.hasOwnProperty("description"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.description);
            if (message.type !== undefined && message.hasOwnProperty("type"))
                writer.uint32(/* id 5, wireType 0 =*/40).uint32(message.type);
            return writer;
        };

        /**
         * Encodes the specified KMDI message, length delimited.
         * @param {kmi.KMDI|Object} message KMDI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMDI.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a KMDI message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMDI} KMDI
         */
        KMDI.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.KMDI();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.version = reader.string();
                    break;
                case 4:
                    message.description = reader.string();
                    break;
                case 5:
                    message.type = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a KMDI message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMDI} KMDI
         */
        KMDI.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a KMDI message.
         * @param {kmi.KMDI|Object} message KMDI message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        KMDI.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.name !== undefined)
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.version !== undefined)
                if (!$util.isString(message.version))
                    return "version: string expected";
            if (message.description !== undefined)
                if (!$util.isString(message.description))
                    return "description: string expected";
            if (message.type !== undefined)
                switch (message.type) {
                default:
                    return "type: enum value expected";
                case 0:
                    break;
                }
            return null;
        };

        /**
         * Creates a KMDI message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMDI} KMDI
         */
        KMDI.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.KMDI)
                return object;
            let message = new $root.kmi.KMDI();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            if (object.name !== undefined && object.name !== null)
                message.name = String(object.name);
            if (object.version !== undefined && object.version !== null)
                message.version = String(object.version);
            if (object.description !== undefined && object.description !== null)
                message.description = String(object.description);
            switch (object.type) {
            case "WEBSERVER":
            case 0:
                message.type = 0;
                break;
            }
            return message;
        };

        /**
         * Creates a KMDI message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMDI.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMDI} KMDI
         */
        KMDI.from = KMDI.fromObject;

        /**
         * Creates a plain object from a KMDI message. Also converts values to other types if specified.
         * @param {kmi.KMDI} message KMDI
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMDI.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = 0;
                object.name = "";
                object.version = "";
                object.description = "";
                object.type = options.enums === String ? "WEBSERVER" : 0;
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.name !== undefined && message.name !== null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.version !== undefined && message.version !== null && message.hasOwnProperty("version"))
                object.version = message.version;
            if (message.description !== undefined && message.description !== null && message.hasOwnProperty("description"))
                object.description = message.description;
            if (message.type !== undefined && message.type !== null && message.hasOwnProperty("type"))
                object.type = options.enums === String ? $types[4][message.type] : message.type;
            return object;
        };

        /**
         * Creates a plain object from this KMDI message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMDI.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this KMDI to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        KMDI.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return KMDI;
    })();

    kmi.FrontendModule = (function() {

        /**
         * Constructs a new FrontendModule.
         * @exports kmi.FrontendModule
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function FrontendModule(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * FrontendModule template.
         * @type {string|undefined}
         */
        FrontendModule.prototype.template = "";

        /**
         * FrontendModule parameters.
         * @type {Object.<string,string>|undefined}
         */
        FrontendModule.prototype.parameters = $util.emptyObject;

        /**
         * Creates a new FrontendModule instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.FrontendModule} FrontendModule instance
         */
        FrontendModule.create = function create(properties) {
            return new FrontendModule(properties);
        };

        /**
         * Encodes the specified FrontendModule message.
         * @param {kmi.FrontendModule|Object} message FrontendModule message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FrontendModule.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.template !== undefined && message.hasOwnProperty("template"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.template);
            if (message.parameters && message.hasOwnProperty("parameters"))
                for (let keys = Object.keys(message.parameters), i = 0; i < keys.length; ++i)
                    writer.uint32(/* id 2, wireType 2 =*/18).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]).uint32(/* id 2, wireType 2 =*/18).string(message.parameters[keys[i]]).ldelim();
            return writer;
        };

        /**
         * Encodes the specified FrontendModule message, length delimited.
         * @param {kmi.FrontendModule|Object} message FrontendModule message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FrontendModule.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a FrontendModule message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.FrontendModule} FrontendModule
         */
        FrontendModule.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.FrontendModule();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.template = reader.string();
                    break;
                case 2:
                    reader.skip().pos++;
                    if (message.parameters === $util.emptyObject)
                        message.parameters = {};
                    let key = reader.string();
                    reader.pos++;
                    message.parameters[typeof key === "object" ? $util.longToHash(key) : key] = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a FrontendModule message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.FrontendModule} FrontendModule
         */
        FrontendModule.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a FrontendModule message.
         * @param {kmi.FrontendModule|Object} message FrontendModule message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        FrontendModule.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.template !== undefined)
                if (!$util.isString(message.template))
                    return "template: string expected";
            if (message.parameters !== undefined) {
                if (!$util.isObject(message.parameters))
                    return "parameters: object expected";
                let key = Object.keys(message.parameters);
                for (let i = 0; i < key.length; ++i)
                    if (!$util.isString(message.parameters[key[i]]))
                        return "parameters: string{k:string} expected";
            }
            return null;
        };

        /**
         * Creates a FrontendModule message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.FrontendModule} FrontendModule
         */
        FrontendModule.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.FrontendModule)
                return object;
            let message = new $root.kmi.FrontendModule();
            if (object.template !== undefined && object.template !== null)
                message.template = String(object.template);
            if (object.parameters) {
                if (typeof object.parameters !== "object")
                    throw TypeError(".kmi.FrontendModule.parameters: object expected");
                message.parameters = {};
                for (let keys = Object.keys(object.parameters), i = 0; i < keys.length; ++i)
                    message.parameters[keys[i]] = String(object.parameters[keys[i]]);
            }
            return message;
        };

        /**
         * Creates a FrontendModule message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.FrontendModule.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.FrontendModule} FrontendModule
         */
        FrontendModule.from = FrontendModule.fromObject;

        /**
         * Creates a plain object from a FrontendModule message. Also converts values to other types if specified.
         * @param {kmi.FrontendModule} message FrontendModule
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        FrontendModule.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.objects || options.defaults)
                object.parameters = {};
            if (options.defaults)
                object.template = "";
            if (message.template !== undefined && message.template !== null && message.hasOwnProperty("template"))
                object.template = message.template;
            if (message.parameters !== undefined && message.parameters !== null && message.hasOwnProperty("parameters")) {
                object.parameters = {};
                for (let keys2 = Object.keys(message.parameters), j = 0; j < keys2.length; ++j)
                    object.parameters[keys2[j]] = message.parameters[keys2[j]];
            }
            return object;
        };

        /**
         * Creates a plain object from this FrontendModule message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        FrontendModule.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this FrontendModule to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        FrontendModule.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return FrontendModule;
    })();

    kmi.KMI = (function() {

        /**
         * Constructs a new KMI.
         * @exports kmi.KMI
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function KMI(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * KMI KMDI.
         * @type {kmi.KMDI|undefined}
         */
        KMI.prototype.KMDI = null;

        /**
         * KMI dockerfile.
         * @type {string|undefined}
         */
        KMI.prototype.dockerfile = "";

        /**
         * KMI container.
         * @type {string|undefined}
         */
        KMI.prototype.container = "";

        /**
         * KMI commands.
         * @type {Object.<string,string>|undefined}
         */
        KMI.prototype.commands = $util.emptyObject;

        /**
         * KMI environment.
         * @type {Object.<string,string>|undefined}
         */
        KMI.prototype.environment = $util.emptyObject;

        /**
         * KMI frontend.
         * @type {Array.<kmi.FrontendModule>|undefined}
         */
        KMI.prototype.frontend = $util.emptyArray;

        /**
         * KMI imports.
         * @type {Array.<string>|undefined}
         */
        KMI.prototype.imports = $util.emptyArray;

        /**
         * KMI interfaces.
         * @type {Object.<string,string>|undefined}
         */
        KMI.prototype.interfaces = $util.emptyObject;

        /**
         * KMI mounts.
         * @type {Array.<string>|undefined}
         */
        KMI.prototype.mounts = $util.emptyArray;

        /**
         * KMI variables.
         * @type {Array.<string>|undefined}
         */
        KMI.prototype.variables = $util.emptyArray;

        /**
         * KMI resources.
         * @type {Object.<string,string>|undefined}
         */
        KMI.prototype.resources = $util.emptyObject;

        // Lazily resolved type references
        const $types = {
            0: "kmi.KMDI",
            5: "kmi.FrontendModule"
        }; $lazyTypes.push($types);

        /**
         * Creates a new KMI instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMI} KMI instance
         */
        KMI.create = function create(properties) {
            return new KMI(properties);
        };

        /**
         * Encodes the specified KMI message.
         * @param {kmi.KMI|Object} message KMI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMI.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.KMDI && message.hasOwnProperty("KMDI"))
                $types[0].encode(message.KMDI, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.dockerfile !== undefined && message.hasOwnProperty("dockerfile"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.dockerfile);
            if (message.container !== undefined && message.hasOwnProperty("container"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.container);
            if (message.commands && message.hasOwnProperty("commands"))
                for (let keys = Object.keys(message.commands), i = 0; i < keys.length; ++i)
                    writer.uint32(/* id 4, wireType 2 =*/34).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]).uint32(/* id 2, wireType 2 =*/18).string(message.commands[keys[i]]).ldelim();
            if (message.environment && message.hasOwnProperty("environment"))
                for (let keys = Object.keys(message.environment), i = 0; i < keys.length; ++i)
                    writer.uint32(/* id 5, wireType 2 =*/42).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]).uint32(/* id 2, wireType 2 =*/18).string(message.environment[keys[i]]).ldelim();
            if (message.frontend !== undefined && message.hasOwnProperty("frontend"))
                for (let i = 0; i < message.frontend.length; ++i)
                    $types[5].encode(message.frontend[i], writer.uint32(/* id 6, wireType 2 =*/50).fork()).ldelim();
            if (message.imports !== undefined && message.hasOwnProperty("imports"))
                for (let i = 0; i < message.imports.length; ++i)
                    writer.uint32(/* id 7, wireType 2 =*/58).string(message.imports[i]);
            if (message.interfaces && message.hasOwnProperty("interfaces"))
                for (let keys = Object.keys(message.interfaces), i = 0; i < keys.length; ++i)
                    writer.uint32(/* id 8, wireType 2 =*/66).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]).uint32(/* id 2, wireType 2 =*/18).string(message.interfaces[keys[i]]).ldelim();
            if (message.mounts !== undefined && message.hasOwnProperty("mounts"))
                for (let i = 0; i < message.mounts.length; ++i)
                    writer.uint32(/* id 9, wireType 2 =*/74).string(message.mounts[i]);
            if (message.variables !== undefined && message.hasOwnProperty("variables"))
                for (let i = 0; i < message.variables.length; ++i)
                    writer.uint32(/* id 10, wireType 2 =*/82).string(message.variables[i]);
            if (message.resources && message.hasOwnProperty("resources"))
                for (let keys = Object.keys(message.resources), i = 0; i < keys.length; ++i)
                    writer.uint32(/* id 11, wireType 2 =*/90).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]).uint32(/* id 2, wireType 2 =*/18).string(message.resources[keys[i]]).ldelim();
            return writer;
        };

        /**
         * Encodes the specified KMI message, length delimited.
         * @param {kmi.KMI|Object} message KMI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMI.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a KMI message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMI} KMI
         */
        KMI.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.KMI();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.KMDI = $types[0].decode(reader, reader.uint32());
                    break;
                case 2:
                    message.dockerfile = reader.string();
                    break;
                case 3:
                    message.container = reader.string();
                    break;
                case 4:
                    reader.skip().pos++;
                    if (message.commands === $util.emptyObject)
                        message.commands = {};
                    let key = reader.string();
                    reader.pos++;
                    message.commands[typeof key === "object" ? $util.longToHash(key) : key] = reader.string();
                    break;
                case 5:
                    reader.skip().pos++;
                    if (message.environment === $util.emptyObject)
                        message.environment = {};
                    let key = reader.string();
                    reader.pos++;
                    message.environment[typeof key === "object" ? $util.longToHash(key) : key] = reader.string();
                    break;
                case 6:
                    if (!(message.frontend && message.frontend.length))
                        message.frontend = [];
                    message.frontend.push($types[5].decode(reader, reader.uint32()));
                    break;
                case 7:
                    if (!(message.imports && message.imports.length))
                        message.imports = [];
                    message.imports.push(reader.string());
                    break;
                case 8:
                    reader.skip().pos++;
                    if (message.interfaces === $util.emptyObject)
                        message.interfaces = {};
                    let key = reader.string();
                    reader.pos++;
                    message.interfaces[typeof key === "object" ? $util.longToHash(key) : key] = reader.string();
                    break;
                case 9:
                    if (!(message.mounts && message.mounts.length))
                        message.mounts = [];
                    message.mounts.push(reader.string());
                    break;
                case 10:
                    if (!(message.variables && message.variables.length))
                        message.variables = [];
                    message.variables.push(reader.string());
                    break;
                case 11:
                    reader.skip().pos++;
                    if (message.resources === $util.emptyObject)
                        message.resources = {};
                    let key = reader.string();
                    reader.pos++;
                    message.resources[typeof key === "object" ? $util.longToHash(key) : key] = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a KMI message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMI} KMI
         */
        KMI.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a KMI message.
         * @param {kmi.KMI|Object} message KMI message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        KMI.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.KMDI !== undefined && message.KMDI !== null) {
                let error = $types[0].verify(message.KMDI);
                if (error)
                    return "KMDI." + error;
            }
            if (message.dockerfile !== undefined)
                if (!$util.isString(message.dockerfile))
                    return "dockerfile: string expected";
            if (message.container !== undefined)
                if (!$util.isString(message.container))
                    return "container: string expected";
            if (message.commands !== undefined) {
                if (!$util.isObject(message.commands))
                    return "commands: object expected";
                let key = Object.keys(message.commands);
                for (let i = 0; i < key.length; ++i)
                    if (!$util.isString(message.commands[key[i]]))
                        return "commands: string{k:string} expected";
            }
            if (message.environment !== undefined) {
                if (!$util.isObject(message.environment))
                    return "environment: object expected";
                let key = Object.keys(message.environment);
                for (let i = 0; i < key.length; ++i)
                    if (!$util.isString(message.environment[key[i]]))
                        return "environment: string{k:string} expected";
            }
            if (message.frontend !== undefined) {
                if (!Array.isArray(message.frontend))
                    return "frontend: array expected";
                for (let i = 0; i < message.frontend.length; ++i) {
                    let error = $types[5].verify(message.frontend[i]);
                    if (error)
                        return "frontend." + error;
                }
            }
            if (message.imports !== undefined) {
                if (!Array.isArray(message.imports))
                    return "imports: array expected";
                for (let i = 0; i < message.imports.length; ++i)
                    if (!$util.isString(message.imports[i]))
                        return "imports: string[] expected";
            }
            if (message.interfaces !== undefined) {
                if (!$util.isObject(message.interfaces))
                    return "interfaces: object expected";
                let key = Object.keys(message.interfaces);
                for (let i = 0; i < key.length; ++i)
                    if (!$util.isString(message.interfaces[key[i]]))
                        return "interfaces: string{k:string} expected";
            }
            if (message.mounts !== undefined) {
                if (!Array.isArray(message.mounts))
                    return "mounts: array expected";
                for (let i = 0; i < message.mounts.length; ++i)
                    if (!$util.isString(message.mounts[i]))
                        return "mounts: string[] expected";
            }
            if (message.variables !== undefined) {
                if (!Array.isArray(message.variables))
                    return "variables: array expected";
                for (let i = 0; i < message.variables.length; ++i)
                    if (!$util.isString(message.variables[i]))
                        return "variables: string[] expected";
            }
            if (message.resources !== undefined) {
                if (!$util.isObject(message.resources))
                    return "resources: object expected";
                let key = Object.keys(message.resources);
                for (let i = 0; i < key.length; ++i)
                    if (!$util.isString(message.resources[key[i]]))
                        return "resources: string{k:string} expected";
            }
            return null;
        };

        /**
         * Creates a KMI message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMI} KMI
         */
        KMI.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.KMI)
                return object;
            let message = new $root.kmi.KMI();
            if (object.KMDI !== undefined && object.KMDI !== null) {
                if (typeof object.KMDI !== "object")
                    throw TypeError(".kmi.KMI.KMDI: object expected");
                message.KMDI = $types[0].fromObject(object.KMDI);
            }
            if (object.dockerfile !== undefined && object.dockerfile !== null)
                message.dockerfile = String(object.dockerfile);
            if (object.container !== undefined && object.container !== null)
                message.container = String(object.container);
            if (object.commands) {
                if (typeof object.commands !== "object")
                    throw TypeError(".kmi.KMI.commands: object expected");
                message.commands = {};
                for (let keys = Object.keys(object.commands), i = 0; i < keys.length; ++i)
                    message.commands[keys[i]] = String(object.commands[keys[i]]);
            }
            if (object.environment) {
                if (typeof object.environment !== "object")
                    throw TypeError(".kmi.KMI.environment: object expected");
                message.environment = {};
                for (let keys = Object.keys(object.environment), i = 0; i < keys.length; ++i)
                    message.environment[keys[i]] = String(object.environment[keys[i]]);
            }
            if (object.frontend) {
                if (!Array.isArray(object.frontend))
                    throw TypeError(".kmi.KMI.frontend: array expected");
                message.frontend = [];
                for (let i = 0; i < object.frontend.length; ++i) {
                    if (typeof object.frontend[i] !== "object")
                        throw TypeError(".kmi.KMI.frontend: object expected");
                    message.frontend[i] = $types[5].fromObject(object.frontend[i]);
                }
            }
            if (object.imports) {
                if (!Array.isArray(object.imports))
                    throw TypeError(".kmi.KMI.imports: array expected");
                message.imports = [];
                for (let i = 0; i < object.imports.length; ++i)
                    message.imports[i] = String(object.imports[i]);
            }
            if (object.interfaces) {
                if (typeof object.interfaces !== "object")
                    throw TypeError(".kmi.KMI.interfaces: object expected");
                message.interfaces = {};
                for (let keys = Object.keys(object.interfaces), i = 0; i < keys.length; ++i)
                    message.interfaces[keys[i]] = String(object.interfaces[keys[i]]);
            }
            if (object.mounts) {
                if (!Array.isArray(object.mounts))
                    throw TypeError(".kmi.KMI.mounts: array expected");
                message.mounts = [];
                for (let i = 0; i < object.mounts.length; ++i)
                    message.mounts[i] = String(object.mounts[i]);
            }
            if (object.variables) {
                if (!Array.isArray(object.variables))
                    throw TypeError(".kmi.KMI.variables: array expected");
                message.variables = [];
                for (let i = 0; i < object.variables.length; ++i)
                    message.variables[i] = String(object.variables[i]);
            }
            if (object.resources) {
                if (typeof object.resources !== "object")
                    throw TypeError(".kmi.KMI.resources: object expected");
                message.resources = {};
                for (let keys = Object.keys(object.resources), i = 0; i < keys.length; ++i)
                    message.resources[keys[i]] = String(object.resources[keys[i]]);
            }
            return message;
        };

        /**
         * Creates a KMI message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMI.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMI} KMI
         */
        KMI.from = KMI.fromObject;

        /**
         * Creates a plain object from a KMI message. Also converts values to other types if specified.
         * @param {kmi.KMI} message KMI
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMI.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults) {
                object.frontend = [];
                object.imports = [];
                object.mounts = [];
                object.variables = [];
            }
            if (options.objects || options.defaults) {
                object.commands = {};
                object.environment = {};
                object.interfaces = {};
                object.resources = {};
            }
            if (options.defaults) {
                object.KMDI = null;
                object.dockerfile = "";
                object.container = "";
            }
            if (message.KMDI !== undefined && message.KMDI !== null && message.hasOwnProperty("KMDI"))
                object.KMDI = $types[0].toObject(message.KMDI, options);
            if (message.dockerfile !== undefined && message.dockerfile !== null && message.hasOwnProperty("dockerfile"))
                object.dockerfile = message.dockerfile;
            if (message.container !== undefined && message.container !== null && message.hasOwnProperty("container"))
                object.container = message.container;
            if (message.commands !== undefined && message.commands !== null && message.hasOwnProperty("commands")) {
                object.commands = {};
                for (let keys2 = Object.keys(message.commands), j = 0; j < keys2.length; ++j)
                    object.commands[keys2[j]] = message.commands[keys2[j]];
            }
            if (message.environment !== undefined && message.environment !== null && message.hasOwnProperty("environment")) {
                object.environment = {};
                for (let keys2 = Object.keys(message.environment), j = 0; j < keys2.length; ++j)
                    object.environment[keys2[j]] = message.environment[keys2[j]];
            }
            if (message.frontend !== undefined && message.frontend !== null && message.hasOwnProperty("frontend")) {
                object.frontend = [];
                for (let j = 0; j < message.frontend.length; ++j)
                    object.frontend[j] = $types[5].toObject(message.frontend[j], options);
            }
            if (message.imports !== undefined && message.imports !== null && message.hasOwnProperty("imports")) {
                object.imports = [];
                for (let j = 0; j < message.imports.length; ++j)
                    object.imports[j] = message.imports[j];
            }
            if (message.interfaces !== undefined && message.interfaces !== null && message.hasOwnProperty("interfaces")) {
                object.interfaces = {};
                for (let keys2 = Object.keys(message.interfaces), j = 0; j < keys2.length; ++j)
                    object.interfaces[keys2[j]] = message.interfaces[keys2[j]];
            }
            if (message.mounts !== undefined && message.mounts !== null && message.hasOwnProperty("mounts")) {
                object.mounts = [];
                for (let j = 0; j < message.mounts.length; ++j)
                    object.mounts[j] = message.mounts[j];
            }
            if (message.variables !== undefined && message.variables !== null && message.hasOwnProperty("variables")) {
                object.variables = [];
                for (let j = 0; j < message.variables.length; ++j)
                    object.variables[j] = message.variables[j];
            }
            if (message.resources !== undefined && message.resources !== null && message.hasOwnProperty("resources")) {
                object.resources = {};
                for (let keys2 = Object.keys(message.resources), j = 0; j < keys2.length; ++j)
                    object.resources[keys2[j]] = message.resources[keys2[j]];
            }
            return object;
        };

        /**
         * Creates a plain object from this KMI message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMI.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this KMI to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        KMI.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return KMI;
    })();

    kmi.AddKMIRequest = (function() {

        /**
         * Constructs a new AddKMIRequest.
         * @exports kmi.AddKMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function AddKMIRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * AddKMIRequest path.
         * @type {string|undefined}
         */
        AddKMIRequest.prototype.path = "";

        /**
         * Creates a new AddKMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.AddKMIRequest} AddKMIRequest instance
         */
        AddKMIRequest.create = function create(properties) {
            return new AddKMIRequest(properties);
        };

        /**
         * Encodes the specified AddKMIRequest message.
         * @param {kmi.AddKMIRequest|Object} message AddKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AddKMIRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.path !== undefined && message.hasOwnProperty("path"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.path);
            return writer;
        };

        /**
         * Encodes the specified AddKMIRequest message, length delimited.
         * @param {kmi.AddKMIRequest|Object} message AddKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AddKMIRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AddKMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        AddKMIRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.AddKMIRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.path = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AddKMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        AddKMIRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AddKMIRequest message.
         * @param {kmi.AddKMIRequest|Object} message AddKMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        AddKMIRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.path !== undefined)
                if (!$util.isString(message.path))
                    return "path: string expected";
            return null;
        };

        /**
         * Creates an AddKMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        AddKMIRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.AddKMIRequest)
                return object;
            let message = new $root.kmi.AddKMIRequest();
            if (object.path !== undefined && object.path !== null)
                message.path = String(object.path);
            return message;
        };

        /**
         * Creates an AddKMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.AddKMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        AddKMIRequest.from = AddKMIRequest.fromObject;

        /**
         * Creates a plain object from an AddKMIRequest message. Also converts values to other types if specified.
         * @param {kmi.AddKMIRequest} message AddKMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AddKMIRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.path = "";
            if (message.path !== undefined && message.path !== null && message.hasOwnProperty("path"))
                object.path = message.path;
            return object;
        };

        /**
         * Creates a plain object from this AddKMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AddKMIRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this AddKMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        AddKMIRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AddKMIRequest;
    })();

    kmi.AddKMIResponse = (function() {

        /**
         * Constructs a new AddKMIResponse.
         * @exports kmi.AddKMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function AddKMIResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * AddKMIResponse ID.
         * @type {number|undefined}
         */
        AddKMIResponse.prototype.ID = 0;

        /**
         * AddKMIResponse error.
         * @type {string|undefined}
         */
        AddKMIResponse.prototype.error = "";

        /**
         * Creates a new AddKMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.AddKMIResponse} AddKMIResponse instance
         */
        AddKMIResponse.create = function create(properties) {
            return new AddKMIResponse(properties);
        };

        /**
         * Encodes the specified AddKMIResponse message.
         * @param {kmi.AddKMIResponse|Object} message AddKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AddKMIResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified AddKMIResponse message, length delimited.
         * @param {kmi.AddKMIResponse|Object} message AddKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        AddKMIResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an AddKMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        AddKMIResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.AddKMIResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an AddKMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        AddKMIResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an AddKMIResponse message.
         * @param {kmi.AddKMIResponse|Object} message AddKMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        AddKMIResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates an AddKMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        AddKMIResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.AddKMIResponse)
                return object;
            let message = new $root.kmi.AddKMIResponse();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates an AddKMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.AddKMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        AddKMIResponse.from = AddKMIResponse.fromObject;

        /**
         * Creates a plain object from an AddKMIResponse message. Also converts values to other types if specified.
         * @param {kmi.AddKMIResponse} message AddKMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AddKMIResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = 0;
                object.error = "";
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this AddKMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        AddKMIResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this AddKMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        AddKMIResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return AddKMIResponse;
    })();

    kmi.RemoveKMIRequest = (function() {

        /**
         * Constructs a new RemoveKMIRequest.
         * @exports kmi.RemoveKMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function RemoveKMIRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * RemoveKMIRequest ID.
         * @type {number|undefined}
         */
        RemoveKMIRequest.prototype.ID = 0;

        /**
         * Creates a new RemoveKMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest instance
         */
        RemoveKMIRequest.create = function create(properties) {
            return new RemoveKMIRequest(properties);
        };

        /**
         * Encodes the specified RemoveKMIRequest message.
         * @param {kmi.RemoveKMIRequest|Object} message RemoveKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveKMIRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            return writer;
        };

        /**
         * Encodes the specified RemoveKMIRequest message, length delimited.
         * @param {kmi.RemoveKMIRequest|Object} message RemoveKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveKMIRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RemoveKMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        RemoveKMIRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.RemoveKMIRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RemoveKMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        RemoveKMIRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RemoveKMIRequest message.
         * @param {kmi.RemoveKMIRequest|Object} message RemoveKMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        RemoveKMIRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            return null;
        };

        /**
         * Creates a RemoveKMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        RemoveKMIRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.RemoveKMIRequest)
                return object;
            let message = new $root.kmi.RemoveKMIRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            return message;
        };

        /**
         * Creates a RemoveKMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.RemoveKMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        RemoveKMIRequest.from = RemoveKMIRequest.fromObject;

        /**
         * Creates a plain object from a RemoveKMIRequest message. Also converts values to other types if specified.
         * @param {kmi.RemoveKMIRequest} message RemoveKMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveKMIRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ID = 0;
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            return object;
        };

        /**
         * Creates a plain object from this RemoveKMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveKMIRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this RemoveKMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        RemoveKMIRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RemoveKMIRequest;
    })();

    kmi.RemoveKMIResponse = (function() {

        /**
         * Constructs a new RemoveKMIResponse.
         * @exports kmi.RemoveKMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function RemoveKMIResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * RemoveKMIResponse error.
         * @type {string|undefined}
         */
        RemoveKMIResponse.prototype.error = "";

        /**
         * Creates a new RemoveKMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse instance
         */
        RemoveKMIResponse.create = function create(properties) {
            return new RemoveKMIResponse(properties);
        };

        /**
         * Encodes the specified RemoveKMIResponse message.
         * @param {kmi.RemoveKMIResponse|Object} message RemoveKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveKMIResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified RemoveKMIResponse message, length delimited.
         * @param {kmi.RemoveKMIResponse|Object} message RemoveKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveKMIResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RemoveKMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        RemoveKMIResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.RemoveKMIResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RemoveKMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        RemoveKMIResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RemoveKMIResponse message.
         * @param {kmi.RemoveKMIResponse|Object} message RemoveKMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        RemoveKMIResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a RemoveKMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        RemoveKMIResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.RemoveKMIResponse)
                return object;
            let message = new $root.kmi.RemoveKMIResponse();
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a RemoveKMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.RemoveKMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        RemoveKMIResponse.from = RemoveKMIResponse.fromObject;

        /**
         * Creates a plain object from a RemoveKMIResponse message. Also converts values to other types if specified.
         * @param {kmi.RemoveKMIResponse} message RemoveKMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveKMIResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.error = "";
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this RemoveKMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveKMIResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this RemoveKMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        RemoveKMIResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RemoveKMIResponse;
    })();

    kmi.GetKMIRequest = (function() {

        /**
         * Constructs a new GetKMIRequest.
         * @exports kmi.GetKMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function GetKMIRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetKMIRequest ID.
         * @type {number|undefined}
         */
        GetKMIRequest.prototype.ID = 0;

        /**
         * Creates a new GetKMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.GetKMIRequest} GetKMIRequest instance
         */
        GetKMIRequest.create = function create(properties) {
            return new GetKMIRequest(properties);
        };

        /**
         * Encodes the specified GetKMIRequest message.
         * @param {kmi.GetKMIRequest|Object} message GetKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetKMIRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            return writer;
        };

        /**
         * Encodes the specified GetKMIRequest message, length delimited.
         * @param {kmi.GetKMIRequest|Object} message GetKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetKMIRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GetKMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        GetKMIRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.GetKMIRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GetKMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        GetKMIRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GetKMIRequest message.
         * @param {kmi.GetKMIRequest|Object} message GetKMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        GetKMIRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            return null;
        };

        /**
         * Creates a GetKMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        GetKMIRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.GetKMIRequest)
                return object;
            let message = new $root.kmi.GetKMIRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = object.ID >>> 0;
            return message;
        };

        /**
         * Creates a GetKMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.GetKMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        GetKMIRequest.from = GetKMIRequest.fromObject;

        /**
         * Creates a plain object from a GetKMIRequest message. Also converts values to other types if specified.
         * @param {kmi.GetKMIRequest} message GetKMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetKMIRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ID = 0;
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            return object;
        };

        /**
         * Creates a plain object from this GetKMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetKMIRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this GetKMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        GetKMIRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return GetKMIRequest;
    })();

    kmi.GetKMIResponse = (function() {

        /**
         * Constructs a new GetKMIResponse.
         * @exports kmi.GetKMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function GetKMIResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * GetKMIResponse kmi.
         * @type {kmi.KMI|undefined}
         */
        GetKMIResponse.prototype.kmi = null;

        /**
         * GetKMIResponse error.
         * @type {string|undefined}
         */
        GetKMIResponse.prototype.error = "";

        // Lazily resolved type references
        const $types = {
            0: "kmi.KMI"
        }; $lazyTypes.push($types);

        /**
         * Creates a new GetKMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.GetKMIResponse} GetKMIResponse instance
         */
        GetKMIResponse.create = function create(properties) {
            return new GetKMIResponse(properties);
        };

        /**
         * Encodes the specified GetKMIResponse message.
         * @param {kmi.GetKMIResponse|Object} message GetKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetKMIResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.kmi && message.hasOwnProperty("kmi"))
                $types[0].encode(message.kmi, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified GetKMIResponse message, length delimited.
         * @param {kmi.GetKMIResponse|Object} message GetKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GetKMIResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GetKMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        GetKMIResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.GetKMIResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.kmi = $types[0].decode(reader, reader.uint32());
                    break;
                case 2:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GetKMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        GetKMIResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GetKMIResponse message.
         * @param {kmi.GetKMIResponse|Object} message GetKMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        GetKMIResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.kmi !== undefined && message.kmi !== null) {
                let error = $types[0].verify(message.kmi);
                if (error)
                    return "kmi." + error;
            }
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a GetKMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        GetKMIResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.GetKMIResponse)
                return object;
            let message = new $root.kmi.GetKMIResponse();
            if (object.kmi !== undefined && object.kmi !== null) {
                if (typeof object.kmi !== "object")
                    throw TypeError(".kmi.GetKMIResponse.kmi: object expected");
                message.kmi = $types[0].fromObject(object.kmi);
            }
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a GetKMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.GetKMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        GetKMIResponse.from = GetKMIResponse.fromObject;

        /**
         * Creates a plain object from a GetKMIResponse message. Also converts values to other types if specified.
         * @param {kmi.GetKMIResponse} message GetKMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetKMIResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.kmi = null;
                object.error = "";
            }
            if (message.kmi !== undefined && message.kmi !== null && message.hasOwnProperty("kmi"))
                object.kmi = $types[0].toObject(message.kmi, options);
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this GetKMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GetKMIResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this GetKMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        GetKMIResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return GetKMIResponse;
    })();

    kmi.KMIRequest = (function() {

        /**
         * Constructs a new KMIRequest.
         * @exports kmi.KMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function KMIRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new KMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMIRequest} KMIRequest instance
         */
        KMIRequest.create = function create(properties) {
            return new KMIRequest(properties);
        };

        /**
         * Encodes the specified KMIRequest message.
         * @param {kmi.KMIRequest|Object} message KMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMIRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified KMIRequest message, length delimited.
         * @param {kmi.KMIRequest|Object} message KMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMIRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a KMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMIRequest} KMIRequest
         */
        KMIRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.KMIRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a KMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMIRequest} KMIRequest
         */
        KMIRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a KMIRequest message.
         * @param {kmi.KMIRequest|Object} message KMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        KMIRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a KMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIRequest} KMIRequest
         */
        KMIRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.KMIRequest)
                return object;
            return new $root.kmi.KMIRequest();
        };

        /**
         * Creates a KMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIRequest} KMIRequest
         */
        KMIRequest.from = KMIRequest.fromObject;

        /**
         * Creates a plain object from a KMIRequest message. Also converts values to other types if specified.
         * @param {kmi.KMIRequest} message KMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMIRequest.toObject = function toObject() {
            return {};
        };

        /**
         * Creates a plain object from this KMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMIRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this KMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        KMIRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return KMIRequest;
    })();

    kmi.KMIResponse = (function() {

        /**
         * Constructs a new KMIResponse.
         * @exports kmi.KMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function KMIResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * KMIResponse kmdi.
         * @type {Array.<kmi.KMDI>|undefined}
         */
        KMIResponse.prototype.kmdi = $util.emptyArray;

        /**
         * KMIResponse error.
         * @type {string|undefined}
         */
        KMIResponse.prototype.error = "";

        // Lazily resolved type references
        const $types = {
            0: "kmi.KMDI"
        }; $lazyTypes.push($types);

        /**
         * Creates a new KMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMIResponse} KMIResponse instance
         */
        KMIResponse.create = function create(properties) {
            return new KMIResponse(properties);
        };

        /**
         * Encodes the specified KMIResponse message.
         * @param {kmi.KMIResponse|Object} message KMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMIResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.kmdi !== undefined && message.hasOwnProperty("kmdi"))
                for (let i = 0; i < message.kmdi.length; ++i)
                    $types[0].encode(message.kmdi[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified KMIResponse message, length delimited.
         * @param {kmi.KMIResponse|Object} message KMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        KMIResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a KMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMIResponse} KMIResponse
         */
        KMIResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.kmi.KMIResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.kmdi && message.kmdi.length))
                        message.kmdi = [];
                    message.kmdi.push($types[0].decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a KMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMIResponse} KMIResponse
         */
        KMIResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a KMIResponse message.
         * @param {kmi.KMIResponse|Object} message KMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        KMIResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.kmdi !== undefined) {
                if (!Array.isArray(message.kmdi))
                    return "kmdi: array expected";
                for (let i = 0; i < message.kmdi.length; ++i) {
                    let error = $types[0].verify(message.kmdi[i]);
                    if (error)
                        return "kmdi." + error;
                }
            }
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a KMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIResponse} KMIResponse
         */
        KMIResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.kmi.KMIResponse)
                return object;
            let message = new $root.kmi.KMIResponse();
            if (object.kmdi) {
                if (!Array.isArray(object.kmdi))
                    throw TypeError(".kmi.KMIResponse.kmdi: array expected");
                message.kmdi = [];
                for (let i = 0; i < object.kmdi.length; ++i) {
                    if (typeof object.kmdi[i] !== "object")
                        throw TypeError(".kmi.KMIResponse.kmdi: object expected");
                    message.kmdi[i] = $types[0].fromObject(object.kmdi[i]);
                }
            }
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a KMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIResponse} KMIResponse
         */
        KMIResponse.from = KMIResponse.fromObject;

        /**
         * Creates a plain object from a KMIResponse message. Also converts values to other types if specified.
         * @param {kmi.KMIResponse} message KMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMIResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.kmdi = [];
            if (options.defaults)
                object.error = "";
            if (message.kmdi !== undefined && message.kmdi !== null && message.hasOwnProperty("kmdi")) {
                object.kmdi = [];
                for (let j = 0; j < message.kmdi.length; ++j)
                    object.kmdi[j] = $types[0].toObject(message.kmdi[j], options);
            }
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this KMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        KMIResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this KMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        KMIResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return KMIResponse;
    })();

    return kmi;
})();

$root.containerlifecycle = (function() {

    /**
     * Namespace containerlifecycle.
     * @exports containerlifecycle
     * @namespace
     */
    const containerlifecycle = {};

    containerlifecycle.ContainerLifecycleService = (function() {

        /**
         * Constructs a new ContainerLifecycleService service.
         * @exports containerlifecycle.ContainerLifecycleService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function ContainerLifecycleService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (ContainerLifecycleService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = ContainerLifecycleService;

        /**
         * Creates new ContainerLifecycleService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {ContainerLifecycleService} RPC service. Useful where requests and/or responses are streamed.
         */
        ContainerLifecycleService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link ContainerLifecycleService#startContainer}.
         * @typedef ContainerLifecycleService_startContainer_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {containerlifecycle.StartContainerResponse} [response] StartContainerResponse
         */

        /**
         * Calls StartContainer.
         * @param {containerlifecycle.StartContainerRequest|Object} request StartContainerRequest message or plain object
         * @param {ContainerLifecycleService_startContainer_Callback} callback Node-style callback called with the error, if any, and StartContainerResponse
         * @returns {undefined}
         */
        ContainerLifecycleService.prototype.startContainer = function startContainer(request, callback) {
            return this.rpcCall(startContainer, $root.containerlifecycle.StartContainerRequest, $root.containerlifecycle.StartContainerResponse, request, callback);
        };

        /**
         * Calls StartContainer.
         * @name ContainerLifecycleService#startContainer
         * @function
         * @param {containerlifecycle.StartContainerRequest|Object} request StartContainerRequest message or plain object
         * @returns {Promise<containerlifecycle.StartContainerResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link ContainerLifecycleService#startCommand}.
         * @typedef ContainerLifecycleService_startCommand_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {containerlifecycle.StartCommandResponse} [response] StartCommandResponse
         */

        /**
         * Calls StartCommand.
         * @param {containerlifecycle.StartCommandRequest|Object} request StartCommandRequest message or plain object
         * @param {ContainerLifecycleService_startCommand_Callback} callback Node-style callback called with the error, if any, and StartCommandResponse
         * @returns {undefined}
         */
        ContainerLifecycleService.prototype.startCommand = function startCommand(request, callback) {
            return this.rpcCall(startCommand, $root.containerlifecycle.StartCommandRequest, $root.containerlifecycle.StartCommandResponse, request, callback);
        };

        /**
         * Calls StartCommand.
         * @name ContainerLifecycleService#startCommand
         * @function
         * @param {containerlifecycle.StartCommandRequest|Object} request StartCommandRequest message or plain object
         * @returns {Promise<containerlifecycle.StartCommandResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link ContainerLifecycleService#stopContainer}.
         * @typedef ContainerLifecycleService_stopContainer_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {containerlifecycle.StopContainerResponse} [response] StopContainerResponse
         */

        /**
         * Calls StopContainer.
         * @param {containerlifecycle.StopContainerRequest|Object} request StopContainerRequest message or plain object
         * @param {ContainerLifecycleService_stopContainer_Callback} callback Node-style callback called with the error, if any, and StopContainerResponse
         * @returns {undefined}
         */
        ContainerLifecycleService.prototype.stopContainer = function stopContainer(request, callback) {
            return this.rpcCall(stopContainer, $root.containerlifecycle.StopContainerRequest, $root.containerlifecycle.StopContainerResponse, request, callback);
        };

        /**
         * Calls StopContainer.
         * @name ContainerLifecycleService#stopContainer
         * @function
         * @param {containerlifecycle.StopContainerRequest|Object} request StopContainerRequest message or plain object
         * @returns {Promise<containerlifecycle.StopContainerResponse>} Promise
         * @variation 2
         */

        return ContainerLifecycleService;
    })();

    containerlifecycle.StartContainerRequest = (function() {

        /**
         * Constructs a new StartContainerRequest.
         * @exports containerlifecycle.StartContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function StartContainerRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * StartContainerRequest ID.
         * @type {string|undefined}
         */
        StartContainerRequest.prototype.ID = "";

        /**
         * Creates a new StartContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest instance
         */
        StartContainerRequest.create = function create(properties) {
            return new StartContainerRequest(properties);
        };

        /**
         * Encodes the specified StartContainerRequest message.
         * @param {containerlifecycle.StartContainerRequest|Object} message StartContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartContainerRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ID);
            return writer;
        };

        /**
         * Encodes the specified StartContainerRequest message, length delimited.
         * @param {containerlifecycle.StartContainerRequest|Object} message StartContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartContainerRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StartContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        StartContainerRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.containerlifecycle.StartContainerRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StartContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        StartContainerRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StartContainerRequest message.
         * @param {containerlifecycle.StartContainerRequest|Object} message StartContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        StartContainerRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            return null;
        };

        /**
         * Creates a StartContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        StartContainerRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.containerlifecycle.StartContainerRequest)
                return object;
            let message = new $root.containerlifecycle.StartContainerRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            return message;
        };

        /**
         * Creates a StartContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        StartContainerRequest.from = StartContainerRequest.fromObject;

        /**
         * Creates a plain object from a StartContainerRequest message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartContainerRequest} message StartContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartContainerRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ID = "";
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            return object;
        };

        /**
         * Creates a plain object from this StartContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartContainerRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this StartContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        StartContainerRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return StartContainerRequest;
    })();

    containerlifecycle.StartCommandRequest = (function() {

        /**
         * Constructs a new StartCommandRequest.
         * @exports containerlifecycle.StartCommandRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function StartCommandRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * StartCommandRequest ID.
         * @type {string|undefined}
         */
        StartCommandRequest.prototype.ID = "";

        /**
         * StartCommandRequest cmd.
         * @type {string|undefined}
         */
        StartCommandRequest.prototype.cmd = "";

        /**
         * Creates a new StartCommandRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest instance
         */
        StartCommandRequest.create = function create(properties) {
            return new StartCommandRequest(properties);
        };

        /**
         * Encodes the specified StartCommandRequest message.
         * @param {containerlifecycle.StartCommandRequest|Object} message StartCommandRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartCommandRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ID);
            if (message.cmd !== undefined && message.hasOwnProperty("cmd"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.cmd);
            return writer;
        };

        /**
         * Encodes the specified StartCommandRequest message, length delimited.
         * @param {containerlifecycle.StartCommandRequest|Object} message StartCommandRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartCommandRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StartCommandRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        StartCommandRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.containerlifecycle.StartCommandRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.string();
                    break;
                case 2:
                    message.cmd = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StartCommandRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        StartCommandRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StartCommandRequest message.
         * @param {containerlifecycle.StartCommandRequest|Object} message StartCommandRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        StartCommandRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            if (message.cmd !== undefined)
                if (!$util.isString(message.cmd))
                    return "cmd: string expected";
            return null;
        };

        /**
         * Creates a StartCommandRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        StartCommandRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.containerlifecycle.StartCommandRequest)
                return object;
            let message = new $root.containerlifecycle.StartCommandRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            if (object.cmd !== undefined && object.cmd !== null)
                message.cmd = String(object.cmd);
            return message;
        };

        /**
         * Creates a StartCommandRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartCommandRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        StartCommandRequest.from = StartCommandRequest.fromObject;

        /**
         * Creates a plain object from a StartCommandRequest message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartCommandRequest} message StartCommandRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartCommandRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = "";
                object.cmd = "";
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.cmd !== undefined && message.cmd !== null && message.hasOwnProperty("cmd"))
                object.cmd = message.cmd;
            return object;
        };

        /**
         * Creates a plain object from this StartCommandRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartCommandRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this StartCommandRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        StartCommandRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return StartCommandRequest;
    })();

    containerlifecycle.StopContainerRequest = (function() {

        /**
         * Constructs a new StopContainerRequest.
         * @exports containerlifecycle.StopContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function StopContainerRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * StopContainerRequest ID.
         * @type {string|undefined}
         */
        StopContainerRequest.prototype.ID = "";

        /**
         * Creates a new StopContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest instance
         */
        StopContainerRequest.create = function create(properties) {
            return new StopContainerRequest(properties);
        };

        /**
         * Encodes the specified StopContainerRequest message.
         * @param {containerlifecycle.StopContainerRequest|Object} message StopContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StopContainerRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ID);
            return writer;
        };

        /**
         * Encodes the specified StopContainerRequest message, length delimited.
         * @param {containerlifecycle.StopContainerRequest|Object} message StopContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StopContainerRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StopContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        StopContainerRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.containerlifecycle.StopContainerRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StopContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        StopContainerRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StopContainerRequest message.
         * @param {containerlifecycle.StopContainerRequest|Object} message StopContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        StopContainerRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            return null;
        };

        /**
         * Creates a StopContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        StopContainerRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.containerlifecycle.StopContainerRequest)
                return object;
            let message = new $root.containerlifecycle.StopContainerRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            return message;
        };

        /**
         * Creates a StopContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StopContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        StopContainerRequest.from = StopContainerRequest.fromObject;

        /**
         * Creates a plain object from a StopContainerRequest message. Also converts values to other types if specified.
         * @param {containerlifecycle.StopContainerRequest} message StopContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StopContainerRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ID = "";
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            return object;
        };

        /**
         * Creates a plain object from this StopContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StopContainerRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this StopContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        StopContainerRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return StopContainerRequest;
    })();

    containerlifecycle.StartContainerResponse = (function() {

        /**
         * Constructs a new StartContainerResponse.
         * @exports containerlifecycle.StartContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function StartContainerResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * StartContainerResponse error.
         * @type {string|undefined}
         */
        StartContainerResponse.prototype.error = "";

        /**
         * Creates a new StartContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse instance
         */
        StartContainerResponse.create = function create(properties) {
            return new StartContainerResponse(properties);
        };

        /**
         * Encodes the specified StartContainerResponse message.
         * @param {containerlifecycle.StartContainerResponse|Object} message StartContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartContainerResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified StartContainerResponse message, length delimited.
         * @param {containerlifecycle.StartContainerResponse|Object} message StartContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartContainerResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StartContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        StartContainerResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.containerlifecycle.StartContainerResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StartContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        StartContainerResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StartContainerResponse message.
         * @param {containerlifecycle.StartContainerResponse|Object} message StartContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        StartContainerResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a StartContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        StartContainerResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.containerlifecycle.StartContainerResponse)
                return object;
            let message = new $root.containerlifecycle.StartContainerResponse();
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a StartContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        StartContainerResponse.from = StartContainerResponse.fromObject;

        /**
         * Creates a plain object from a StartContainerResponse message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartContainerResponse} message StartContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartContainerResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.error = "";
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this StartContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartContainerResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this StartContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        StartContainerResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return StartContainerResponse;
    })();

    containerlifecycle.StartCommandResponse = (function() {

        /**
         * Constructs a new StartCommandResponse.
         * @exports containerlifecycle.StartCommandResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function StartCommandResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * StartCommandResponse ID.
         * @type {string|undefined}
         */
        StartCommandResponse.prototype.ID = "";

        /**
         * StartCommandResponse error.
         * @type {string|undefined}
         */
        StartCommandResponse.prototype.error = "";

        /**
         * Creates a new StartCommandResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse instance
         */
        StartCommandResponse.create = function create(properties) {
            return new StartCommandResponse(properties);
        };

        /**
         * Encodes the specified StartCommandResponse message.
         * @param {containerlifecycle.StartCommandResponse|Object} message StartCommandResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartCommandResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ID);
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified StartCommandResponse message, length delimited.
         * @param {containerlifecycle.StartCommandResponse|Object} message StartCommandResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StartCommandResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StartCommandResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        StartCommandResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.containerlifecycle.StartCommandResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.string();
                    break;
                case 2:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StartCommandResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        StartCommandResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StartCommandResponse message.
         * @param {containerlifecycle.StartCommandResponse|Object} message StartCommandResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        StartCommandResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a StartCommandResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        StartCommandResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.containerlifecycle.StartCommandResponse)
                return object;
            let message = new $root.containerlifecycle.StartCommandResponse();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a StartCommandResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartCommandResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        StartCommandResponse.from = StartCommandResponse.fromObject;

        /**
         * Creates a plain object from a StartCommandResponse message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartCommandResponse} message StartCommandResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartCommandResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = "";
                object.error = "";
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this StartCommandResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StartCommandResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this StartCommandResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        StartCommandResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return StartCommandResponse;
    })();

    containerlifecycle.StopContainerResponse = (function() {

        /**
         * Constructs a new StopContainerResponse.
         * @exports containerlifecycle.StopContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function StopContainerResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * StopContainerResponse error.
         * @type {string|undefined}
         */
        StopContainerResponse.prototype.error = "";

        /**
         * Creates a new StopContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse instance
         */
        StopContainerResponse.create = function create(properties) {
            return new StopContainerResponse(properties);
        };

        /**
         * Encodes the specified StopContainerResponse message.
         * @param {containerlifecycle.StopContainerResponse|Object} message StopContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StopContainerResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.error !== undefined && message.hasOwnProperty("error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.error);
            return writer;
        };

        /**
         * Encodes the specified StopContainerResponse message, length delimited.
         * @param {containerlifecycle.StopContainerResponse|Object} message StopContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        StopContainerResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a StopContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        StopContainerResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.containerlifecycle.StopContainerResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a StopContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        StopContainerResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a StopContainerResponse message.
         * @param {containerlifecycle.StopContainerResponse|Object} message StopContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        StopContainerResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.error !== undefined)
                if (!$util.isString(message.error))
                    return "error: string expected";
            return null;
        };

        /**
         * Creates a StopContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        StopContainerResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.containerlifecycle.StopContainerResponse)
                return object;
            let message = new $root.containerlifecycle.StopContainerResponse();
            if (object.error !== undefined && object.error !== null)
                message.error = String(object.error);
            return message;
        };

        /**
         * Creates a StopContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StopContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        StopContainerResponse.from = StopContainerResponse.fromObject;

        /**
         * Creates a plain object from a StopContainerResponse message. Also converts values to other types if specified.
         * @param {containerlifecycle.StopContainerResponse} message StopContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StopContainerResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.error = "";
            if (message.error !== undefined && message.error !== null && message.hasOwnProperty("error"))
                object.error = message.error;
            return object;
        };

        /**
         * Creates a plain object from this StopContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        StopContainerResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this StopContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        StopContainerResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return StopContainerResponse;
    })();

    return containerlifecycle;
})();

$root.customercontainer = (function() {

    /**
     * Namespace customercontainer.
     * @exports customercontainer
     * @namespace
     */
    const customercontainer = {};

    customercontainer.CustomerContainerService = (function() {

        /**
         * Constructs a new CustomerContainerService service.
         * @exports customercontainer.CustomerContainerService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function CustomerContainerService(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (CustomerContainerService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = CustomerContainerService;

        /**
         * Creates new CustomerContainerService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {CustomerContainerService} RPC service. Useful where requests and/or responses are streamed.
         */
        CustomerContainerService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link CustomerContainerService#createContainer}.
         * @typedef CustomerContainerService_createContainer_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {customercontainer.CreateContainerResponse} [response] CreateContainerResponse
         */

        /**
         * Calls CreateContainer.
         * @param {customercontainer.CreateContainerRequest|Object} request CreateContainerRequest message or plain object
         * @param {CustomerContainerService_createContainer_Callback} callback Node-style callback called with the error, if any, and CreateContainerResponse
         * @returns {undefined}
         */
        CustomerContainerService.prototype.createContainer = function createContainer(request, callback) {
            return this.rpcCall(createContainer, $root.customercontainer.CreateContainerRequest, $root.customercontainer.CreateContainerResponse, request, callback);
        };

        /**
         * Calls CreateContainer.
         * @name CustomerContainerService#createContainer
         * @function
         * @param {customercontainer.CreateContainerRequest|Object} request CreateContainerRequest message or plain object
         * @returns {Promise<customercontainer.CreateContainerResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link CustomerContainerService#editContainer}.
         * @typedef CustomerContainerService_editContainer_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {customercontainer.EditContainerResponse} [response] EditContainerResponse
         */

        /**
         * Calls EditContainer.
         * @param {customercontainer.EditContainerRequest|Object} request EditContainerRequest message or plain object
         * @param {CustomerContainerService_editContainer_Callback} callback Node-style callback called with the error, if any, and EditContainerResponse
         * @returns {undefined}
         */
        CustomerContainerService.prototype.editContainer = function editContainer(request, callback) {
            return this.rpcCall(editContainer, $root.customercontainer.EditContainerRequest, $root.customercontainer.EditContainerResponse, request, callback);
        };

        /**
         * Calls EditContainer.
         * @name CustomerContainerService#editContainer
         * @function
         * @param {customercontainer.EditContainerRequest|Object} request EditContainerRequest message or plain object
         * @returns {Promise<customercontainer.EditContainerResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link CustomerContainerService#removeContainer}.
         * @typedef CustomerContainerService_removeContainer_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {customercontainer.RemoveContainerResponse} [response] RemoveContainerResponse
         */

        /**
         * Calls RemoveContainer.
         * @param {customercontainer.RemoveContainerRequest|Object} request RemoveContainerRequest message or plain object
         * @param {CustomerContainerService_removeContainer_Callback} callback Node-style callback called with the error, if any, and RemoveContainerResponse
         * @returns {undefined}
         */
        CustomerContainerService.prototype.removeContainer = function removeContainer(request, callback) {
            return this.rpcCall(removeContainer, $root.customercontainer.RemoveContainerRequest, $root.customercontainer.RemoveContainerResponse, request, callback);
        };

        /**
         * Calls RemoveContainer.
         * @name CustomerContainerService#removeContainer
         * @function
         * @param {customercontainer.RemoveContainerRequest|Object} request RemoveContainerRequest message or plain object
         * @returns {Promise<customercontainer.RemoveContainerResponse>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link CustomerContainerService#instances}.
         * @typedef CustomerContainerService_instances_Callback
         * @type {function}
         * @param {?Error} error Error, if any
         * @param {customercontainer.InstancesResponse} [response] InstancesResponse
         */

        /**
         * Calls Instances.
         * @param {customercontainer.InstancesRequest|Object} request InstancesRequest message or plain object
         * @param {CustomerContainerService_instances_Callback} callback Node-style callback called with the error, if any, and InstancesResponse
         * @returns {undefined}
         */
        CustomerContainerService.prototype.instances = function instances(request, callback) {
            return this.rpcCall(instances, $root.customercontainer.InstancesRequest, $root.customercontainer.InstancesResponse, request, callback);
        };

        /**
         * Calls Instances.
         * @name CustomerContainerService#instances
         * @function
         * @param {customercontainer.InstancesRequest|Object} request InstancesRequest message or plain object
         * @returns {Promise<customercontainer.InstancesResponse>} Promise
         * @variation 2
         */

        return CustomerContainerService;
    })();

    customercontainer.ContainerConfig = (function() {

        /**
         * Constructs a new ContainerConfig.
         * @exports customercontainer.ContainerConfig
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function ContainerConfig(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * ContainerConfig ImageName.
         * @type {string|undefined}
         */
        ContainerConfig.prototype.ImageName = "";

        /**
         * Creates a new ContainerConfig instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.ContainerConfig} ContainerConfig instance
         */
        ContainerConfig.create = function create(properties) {
            return new ContainerConfig(properties);
        };

        /**
         * Encodes the specified ContainerConfig message.
         * @param {customercontainer.ContainerConfig|Object} message ContainerConfig message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ContainerConfig.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ImageName !== undefined && message.hasOwnProperty("ImageName"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ImageName);
            return writer;
        };

        /**
         * Encodes the specified ContainerConfig message, length delimited.
         * @param {customercontainer.ContainerConfig|Object} message ContainerConfig message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ContainerConfig.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ContainerConfig message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        ContainerConfig.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.ContainerConfig();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ImageName = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ContainerConfig message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        ContainerConfig.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ContainerConfig message.
         * @param {customercontainer.ContainerConfig|Object} message ContainerConfig message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        ContainerConfig.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ImageName !== undefined)
                if (!$util.isString(message.ImageName))
                    return "ImageName: string expected";
            return null;
        };

        /**
         * Creates a ContainerConfig message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        ContainerConfig.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.ContainerConfig)
                return object;
            let message = new $root.customercontainer.ContainerConfig();
            if (object.ImageName !== undefined && object.ImageName !== null)
                message.ImageName = String(object.ImageName);
            return message;
        };

        /**
         * Creates a ContainerConfig message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.ContainerConfig.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        ContainerConfig.from = ContainerConfig.fromObject;

        /**
         * Creates a plain object from a ContainerConfig message. Also converts values to other types if specified.
         * @param {customercontainer.ContainerConfig} message ContainerConfig
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ContainerConfig.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ImageName = "";
            if (message.ImageName !== undefined && message.ImageName !== null && message.hasOwnProperty("ImageName"))
                object.ImageName = message.ImageName;
            return object;
        };

        /**
         * Creates a plain object from this ContainerConfig message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ContainerConfig.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this ContainerConfig to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        ContainerConfig.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ContainerConfig;
    })();

    customercontainer.CreateContainerRequest = (function() {

        /**
         * Constructs a new CreateContainerRequest.
         * @exports customercontainer.CreateContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function CreateContainerRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * CreateContainerRequest Refid.
         * @type {number|undefined}
         */
        CreateContainerRequest.prototype.Refid = 0;

        /**
         * CreateContainerRequest Cfg.
         * @type {customercontainer.ContainerConfig|undefined}
         */
        CreateContainerRequest.prototype.Cfg = null;

        // Lazily resolved type references
        const $types = {
            1: "customercontainer.ContainerConfig"
        }; $lazyTypes.push($types);

        /**
         * Creates a new CreateContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest instance
         */
        CreateContainerRequest.create = function create(properties) {
            return new CreateContainerRequest(properties);
        };

        /**
         * Encodes the specified CreateContainerRequest message.
         * @param {customercontainer.CreateContainerRequest|Object} message CreateContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateContainerRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Refid !== undefined && message.hasOwnProperty("Refid"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.Refid);
            if (message.Cfg && message.hasOwnProperty("Cfg"))
                $types[1].encode(message.Cfg, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified CreateContainerRequest message, length delimited.
         * @param {customercontainer.CreateContainerRequest|Object} message CreateContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateContainerRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CreateContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        CreateContainerRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.CreateContainerRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Refid = reader.uint32();
                    break;
                case 2:
                    message.Cfg = $types[1].decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CreateContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        CreateContainerRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CreateContainerRequest message.
         * @param {customercontainer.CreateContainerRequest|Object} message CreateContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        CreateContainerRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Refid !== undefined)
                if (!$util.isInteger(message.Refid))
                    return "Refid: integer expected";
            if (message.Cfg !== undefined && message.Cfg !== null) {
                let error = $types[1].verify(message.Cfg);
                if (error)
                    return "Cfg." + error;
            }
            return null;
        };

        /**
         * Creates a CreateContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        CreateContainerRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.CreateContainerRequest)
                return object;
            let message = new $root.customercontainer.CreateContainerRequest();
            if (object.Refid !== undefined && object.Refid !== null)
                message.Refid = object.Refid >>> 0;
            if (object.Cfg !== undefined && object.Cfg !== null) {
                if (typeof object.Cfg !== "object")
                    throw TypeError(".customercontainer.CreateContainerRequest.Cfg: object expected");
                message.Cfg = $types[1].fromObject(object.Cfg);
            }
            return message;
        };

        /**
         * Creates a CreateContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        CreateContainerRequest.from = CreateContainerRequest.fromObject;

        /**
         * Creates a plain object from a CreateContainerRequest message. Also converts values to other types if specified.
         * @param {customercontainer.CreateContainerRequest} message CreateContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateContainerRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.Refid = 0;
                object.Cfg = null;
            }
            if (message.Refid !== undefined && message.Refid !== null && message.hasOwnProperty("Refid"))
                object.Refid = message.Refid;
            if (message.Cfg !== undefined && message.Cfg !== null && message.hasOwnProperty("Cfg"))
                object.Cfg = $types[1].toObject(message.Cfg, options);
            return object;
        };

        /**
         * Creates a plain object from this CreateContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateContainerRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this CreateContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        CreateContainerRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CreateContainerRequest;
    })();

    customercontainer.CreateContainerResponse = (function() {

        /**
         * Constructs a new CreateContainerResponse.
         * @exports customercontainer.CreateContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function CreateContainerResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * CreateContainerResponse Name.
         * @type {string|undefined}
         */
        CreateContainerResponse.prototype.Name = "";

        /**
         * CreateContainerResponse ID.
         * @type {string|undefined}
         */
        CreateContainerResponse.prototype.ID = "";

        /**
         * CreateContainerResponse Error.
         * @type {string|undefined}
         */
        CreateContainerResponse.prototype.Error = "";

        /**
         * Creates a new CreateContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse instance
         */
        CreateContainerResponse.create = function create(properties) {
            return new CreateContainerResponse(properties);
        };

        /**
         * Encodes the specified CreateContainerResponse message.
         * @param {customercontainer.CreateContainerResponse|Object} message CreateContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateContainerResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Name !== undefined && message.hasOwnProperty("Name"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.Name);
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.ID);
            if (message.Error !== undefined && message.hasOwnProperty("Error"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.Error);
            return writer;
        };

        /**
         * Encodes the specified CreateContainerResponse message, length delimited.
         * @param {customercontainer.CreateContainerResponse|Object} message CreateContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateContainerResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CreateContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        CreateContainerResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.CreateContainerResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Name = reader.string();
                    break;
                case 2:
                    message.ID = reader.string();
                    break;
                case 3:
                    message.Error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CreateContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        CreateContainerResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CreateContainerResponse message.
         * @param {customercontainer.CreateContainerResponse|Object} message CreateContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        CreateContainerResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Name !== undefined)
                if (!$util.isString(message.Name))
                    return "Name: string expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            if (message.Error !== undefined)
                if (!$util.isString(message.Error))
                    return "Error: string expected";
            return null;
        };

        /**
         * Creates a CreateContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        CreateContainerResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.CreateContainerResponse)
                return object;
            let message = new $root.customercontainer.CreateContainerResponse();
            if (object.Name !== undefined && object.Name !== null)
                message.Name = String(object.Name);
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            if (object.Error !== undefined && object.Error !== null)
                message.Error = String(object.Error);
            return message;
        };

        /**
         * Creates a CreateContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        CreateContainerResponse.from = CreateContainerResponse.fromObject;

        /**
         * Creates a plain object from a CreateContainerResponse message. Also converts values to other types if specified.
         * @param {customercontainer.CreateContainerResponse} message CreateContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateContainerResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.Name = "";
                object.ID = "";
                object.Error = "";
            }
            if (message.Name !== undefined && message.Name !== null && message.hasOwnProperty("Name"))
                object.Name = message.Name;
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.Error !== undefined && message.Error !== null && message.hasOwnProperty("Error"))
                object.Error = message.Error;
            return object;
        };

        /**
         * Creates a plain object from this CreateContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateContainerResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this CreateContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        CreateContainerResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CreateContainerResponse;
    })();

    customercontainer.EditContainerRequest = (function() {

        /**
         * Constructs a new EditContainerRequest.
         * @exports customercontainer.EditContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function EditContainerRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * EditContainerRequest ID.
         * @type {string|undefined}
         */
        EditContainerRequest.prototype.ID = "";

        /**
         * EditContainerRequest Cfg.
         * @type {customercontainer.ContainerConfig|undefined}
         */
        EditContainerRequest.prototype.Cfg = null;

        // Lazily resolved type references
        const $types = {
            1: "customercontainer.ContainerConfig"
        }; $lazyTypes.push($types);

        /**
         * Creates a new EditContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest instance
         */
        EditContainerRequest.create = function create(properties) {
            return new EditContainerRequest(properties);
        };

        /**
         * Encodes the specified EditContainerRequest message.
         * @param {customercontainer.EditContainerRequest|Object} message EditContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditContainerRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ID);
            if (message.Cfg && message.hasOwnProperty("Cfg"))
                $types[1].encode(message.Cfg, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified EditContainerRequest message, length delimited.
         * @param {customercontainer.EditContainerRequest|Object} message EditContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditContainerRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an EditContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        EditContainerRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.EditContainerRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.string();
                    break;
                case 2:
                    message.Cfg = $types[1].decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an EditContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        EditContainerRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an EditContainerRequest message.
         * @param {customercontainer.EditContainerRequest|Object} message EditContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        EditContainerRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            if (message.Cfg !== undefined && message.Cfg !== null) {
                let error = $types[1].verify(message.Cfg);
                if (error)
                    return "Cfg." + error;
            }
            return null;
        };

        /**
         * Creates an EditContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        EditContainerRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.EditContainerRequest)
                return object;
            let message = new $root.customercontainer.EditContainerRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            if (object.Cfg !== undefined && object.Cfg !== null) {
                if (typeof object.Cfg !== "object")
                    throw TypeError(".customercontainer.EditContainerRequest.Cfg: object expected");
                message.Cfg = $types[1].fromObject(object.Cfg);
            }
            return message;
        };

        /**
         * Creates an EditContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.EditContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        EditContainerRequest.from = EditContainerRequest.fromObject;

        /**
         * Creates a plain object from an EditContainerRequest message. Also converts values to other types if specified.
         * @param {customercontainer.EditContainerRequest} message EditContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditContainerRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = "";
                object.Cfg = null;
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.Cfg !== undefined && message.Cfg !== null && message.hasOwnProperty("Cfg"))
                object.Cfg = $types[1].toObject(message.Cfg, options);
            return object;
        };

        /**
         * Creates a plain object from this EditContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditContainerRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this EditContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        EditContainerRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return EditContainerRequest;
    })();

    customercontainer.EditContainerResponse = (function() {

        /**
         * Constructs a new EditContainerResponse.
         * @exports customercontainer.EditContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function EditContainerResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * EditContainerResponse Error.
         * @type {string|undefined}
         */
        EditContainerResponse.prototype.Error = "";

        /**
         * Creates a new EditContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse instance
         */
        EditContainerResponse.create = function create(properties) {
            return new EditContainerResponse(properties);
        };

        /**
         * Encodes the specified EditContainerResponse message.
         * @param {customercontainer.EditContainerResponse|Object} message EditContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditContainerResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Error !== undefined && message.hasOwnProperty("Error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.Error);
            return writer;
        };

        /**
         * Encodes the specified EditContainerResponse message, length delimited.
         * @param {customercontainer.EditContainerResponse|Object} message EditContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EditContainerResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an EditContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        EditContainerResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.EditContainerResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an EditContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        EditContainerResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an EditContainerResponse message.
         * @param {customercontainer.EditContainerResponse|Object} message EditContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        EditContainerResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Error !== undefined)
                if (!$util.isString(message.Error))
                    return "Error: string expected";
            return null;
        };

        /**
         * Creates an EditContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        EditContainerResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.EditContainerResponse)
                return object;
            let message = new $root.customercontainer.EditContainerResponse();
            if (object.Error !== undefined && object.Error !== null)
                message.Error = String(object.Error);
            return message;
        };

        /**
         * Creates an EditContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.EditContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        EditContainerResponse.from = EditContainerResponse.fromObject;

        /**
         * Creates a plain object from an EditContainerResponse message. Also converts values to other types if specified.
         * @param {customercontainer.EditContainerResponse} message EditContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditContainerResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.Error = "";
            if (message.Error !== undefined && message.Error !== null && message.hasOwnProperty("Error"))
                object.Error = message.Error;
            return object;
        };

        /**
         * Creates a plain object from this EditContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EditContainerResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this EditContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        EditContainerResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return EditContainerResponse;
    })();

    customercontainer.RemoveContainerRequest = (function() {

        /**
         * Constructs a new RemoveContainerRequest.
         * @exports customercontainer.RemoveContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function RemoveContainerRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * RemoveContainerRequest ID.
         * @type {string|undefined}
         */
        RemoveContainerRequest.prototype.ID = "";

        /**
         * Creates a new RemoveContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest instance
         */
        RemoveContainerRequest.create = function create(properties) {
            return new RemoveContainerRequest(properties);
        };

        /**
         * Encodes the specified RemoveContainerRequest message.
         * @param {customercontainer.RemoveContainerRequest|Object} message RemoveContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveContainerRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ID);
            return writer;
        };

        /**
         * Encodes the specified RemoveContainerRequest message, length delimited.
         * @param {customercontainer.RemoveContainerRequest|Object} message RemoveContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveContainerRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RemoveContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        RemoveContainerRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.RemoveContainerRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RemoveContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        RemoveContainerRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RemoveContainerRequest message.
         * @param {customercontainer.RemoveContainerRequest|Object} message RemoveContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        RemoveContainerRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            return null;
        };

        /**
         * Creates a RemoveContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        RemoveContainerRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.RemoveContainerRequest)
                return object;
            let message = new $root.customercontainer.RemoveContainerRequest();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            return message;
        };

        /**
         * Creates a RemoveContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.RemoveContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        RemoveContainerRequest.from = RemoveContainerRequest.fromObject;

        /**
         * Creates a plain object from a RemoveContainerRequest message. Also converts values to other types if specified.
         * @param {customercontainer.RemoveContainerRequest} message RemoveContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveContainerRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.ID = "";
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            return object;
        };

        /**
         * Creates a plain object from this RemoveContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveContainerRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this RemoveContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        RemoveContainerRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RemoveContainerRequest;
    })();

    customercontainer.RemoveContainerResponse = (function() {

        /**
         * Constructs a new RemoveContainerResponse.
         * @exports customercontainer.RemoveContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function RemoveContainerResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * RemoveContainerResponse Error.
         * @type {string|undefined}
         */
        RemoveContainerResponse.prototype.Error = "";

        /**
         * Creates a new RemoveContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse instance
         */
        RemoveContainerResponse.create = function create(properties) {
            return new RemoveContainerResponse(properties);
        };

        /**
         * Encodes the specified RemoveContainerResponse message.
         * @param {customercontainer.RemoveContainerResponse|Object} message RemoveContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveContainerResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Error !== undefined && message.hasOwnProperty("Error"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.Error);
            return writer;
        };

        /**
         * Encodes the specified RemoveContainerResponse message, length delimited.
         * @param {customercontainer.RemoveContainerResponse|Object} message RemoveContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RemoveContainerResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RemoveContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        RemoveContainerResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.RemoveContainerResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RemoveContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        RemoveContainerResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RemoveContainerResponse message.
         * @param {customercontainer.RemoveContainerResponse|Object} message RemoveContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        RemoveContainerResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Error !== undefined)
                if (!$util.isString(message.Error))
                    return "Error: string expected";
            return null;
        };

        /**
         * Creates a RemoveContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        RemoveContainerResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.RemoveContainerResponse)
                return object;
            let message = new $root.customercontainer.RemoveContainerResponse();
            if (object.Error !== undefined && object.Error !== null)
                message.Error = String(object.Error);
            return message;
        };

        /**
         * Creates a RemoveContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.RemoveContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        RemoveContainerResponse.from = RemoveContainerResponse.fromObject;

        /**
         * Creates a plain object from a RemoveContainerResponse message. Also converts values to other types if specified.
         * @param {customercontainer.RemoveContainerResponse} message RemoveContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveContainerResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.Error = "";
            if (message.Error !== undefined && message.Error !== null && message.hasOwnProperty("Error"))
                object.Error = message.Error;
            return object;
        };

        /**
         * Creates a plain object from this RemoveContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RemoveContainerResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this RemoveContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        RemoveContainerResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RemoveContainerResponse;
    })();

    customercontainer.InstancesRequest = (function() {

        /**
         * Constructs a new InstancesRequest.
         * @exports customercontainer.InstancesRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function InstancesRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * InstancesRequest Refid.
         * @type {number|undefined}
         */
        InstancesRequest.prototype.Refid = 0;

        /**
         * Creates a new InstancesRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.InstancesRequest} InstancesRequest instance
         */
        InstancesRequest.create = function create(properties) {
            return new InstancesRequest(properties);
        };

        /**
         * Encodes the specified InstancesRequest message.
         * @param {customercontainer.InstancesRequest|Object} message InstancesRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InstancesRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Refid !== undefined && message.hasOwnProperty("Refid"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.Refid);
            return writer;
        };

        /**
         * Encodes the specified InstancesRequest message, length delimited.
         * @param {customercontainer.InstancesRequest|Object} message InstancesRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InstancesRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an InstancesRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        InstancesRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.InstancesRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Refid = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an InstancesRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        InstancesRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an InstancesRequest message.
         * @param {customercontainer.InstancesRequest|Object} message InstancesRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        InstancesRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Refid !== undefined)
                if (!$util.isInteger(message.Refid))
                    return "Refid: integer expected";
            return null;
        };

        /**
         * Creates an InstancesRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        InstancesRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.InstancesRequest)
                return object;
            let message = new $root.customercontainer.InstancesRequest();
            if (object.Refid !== undefined && object.Refid !== null)
                message.Refid = object.Refid >>> 0;
            return message;
        };

        /**
         * Creates an InstancesRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.InstancesRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        InstancesRequest.from = InstancesRequest.fromObject;

        /**
         * Creates a plain object from an InstancesRequest message. Also converts values to other types if specified.
         * @param {customercontainer.InstancesRequest} message InstancesRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InstancesRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults)
                object.Refid = 0;
            if (message.Refid !== undefined && message.Refid !== null && message.hasOwnProperty("Refid"))
                object.Refid = message.Refid;
            return object;
        };

        /**
         * Creates a plain object from this InstancesRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InstancesRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this InstancesRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        InstancesRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return InstancesRequest;
    })();

    customercontainer.InstancesResponse = (function() {

        /**
         * Constructs a new InstancesResponse.
         * @exports customercontainer.InstancesResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function InstancesResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * InstancesResponse Instances.
         * @type {Array.<string>|undefined}
         */
        InstancesResponse.prototype.Instances = $util.emptyArray;

        /**
         * Creates a new InstancesResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.InstancesResponse} InstancesResponse instance
         */
        InstancesResponse.create = function create(properties) {
            return new InstancesResponse(properties);
        };

        /**
         * Encodes the specified InstancesResponse message.
         * @param {customercontainer.InstancesResponse|Object} message InstancesResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InstancesResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Instances !== undefined && message.hasOwnProperty("Instances"))
                for (let i = 0; i < message.Instances.length; ++i)
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.Instances[i]);
            return writer;
        };

        /**
         * Encodes the specified InstancesResponse message, length delimited.
         * @param {customercontainer.InstancesResponse|Object} message InstancesResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InstancesResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an InstancesResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        InstancesResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.InstancesResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.Instances && message.Instances.length))
                        message.Instances = [];
                    message.Instances.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an InstancesResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        InstancesResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an InstancesResponse message.
         * @param {customercontainer.InstancesResponse|Object} message InstancesResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        InstancesResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Instances !== undefined) {
                if (!Array.isArray(message.Instances))
                    return "Instances: array expected";
                for (let i = 0; i < message.Instances.length; ++i)
                    if (!$util.isString(message.Instances[i]))
                        return "Instances: string[] expected";
            }
            return null;
        };

        /**
         * Creates an InstancesResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        InstancesResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.InstancesResponse)
                return object;
            let message = new $root.customercontainer.InstancesResponse();
            if (object.Instances) {
                if (!Array.isArray(object.Instances))
                    throw TypeError(".customercontainer.InstancesResponse.Instances: array expected");
                message.Instances = [];
                for (let i = 0; i < object.Instances.length; ++i)
                    message.Instances[i] = String(object.Instances[i]);
            }
            return message;
        };

        /**
         * Creates an InstancesResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.InstancesResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        InstancesResponse.from = InstancesResponse.fromObject;

        /**
         * Creates a plain object from an InstancesResponse message. Also converts values to other types if specified.
         * @param {customercontainer.InstancesResponse} message InstancesResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InstancesResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.arrays || options.defaults)
                object.Instances = [];
            if (message.Instances !== undefined && message.Instances !== null && message.hasOwnProperty("Instances")) {
                object.Instances = [];
                for (let j = 0; j < message.Instances.length; ++j)
                    object.Instances[j] = message.Instances[j];
            }
            return object;
        };

        /**
         * Creates a plain object from this InstancesResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InstancesResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this InstancesResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        InstancesResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return InstancesResponse;
    })();

    customercontainer.CreateDockerImageRequest = (function() {

        /**
         * Constructs a new CreateDockerImageRequest.
         * @exports customercontainer.CreateDockerImageRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function CreateDockerImageRequest(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * CreateDockerImageRequest Refid.
         * @type {number|undefined}
         */
        CreateDockerImageRequest.prototype.Refid = 0;

        /**
         * CreateDockerImageRequest KmiID.
         * @type {number|undefined}
         */
        CreateDockerImageRequest.prototype.KmiID = 0;

        /**
         * Creates a new CreateDockerImageRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest instance
         */
        CreateDockerImageRequest.create = function create(properties) {
            return new CreateDockerImageRequest(properties);
        };

        /**
         * Encodes the specified CreateDockerImageRequest message.
         * @param {customercontainer.CreateDockerImageRequest|Object} message CreateDockerImageRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateDockerImageRequest.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Refid !== undefined && message.hasOwnProperty("Refid"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.Refid);
            if (message.KmiID !== undefined && message.hasOwnProperty("KmiID"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint32(message.KmiID);
            return writer;
        };

        /**
         * Encodes the specified CreateDockerImageRequest message, length delimited.
         * @param {customercontainer.CreateDockerImageRequest|Object} message CreateDockerImageRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateDockerImageRequest.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CreateDockerImageRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        CreateDockerImageRequest.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.CreateDockerImageRequest();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Refid = reader.uint32();
                    break;
                case 2:
                    message.KmiID = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CreateDockerImageRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        CreateDockerImageRequest.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CreateDockerImageRequest message.
         * @param {customercontainer.CreateDockerImageRequest|Object} message CreateDockerImageRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        CreateDockerImageRequest.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Refid !== undefined)
                if (!$util.isInteger(message.Refid))
                    return "Refid: integer expected";
            if (message.KmiID !== undefined)
                if (!$util.isInteger(message.KmiID))
                    return "KmiID: integer expected";
            return null;
        };

        /**
         * Creates a CreateDockerImageRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        CreateDockerImageRequest.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.CreateDockerImageRequest)
                return object;
            let message = new $root.customercontainer.CreateDockerImageRequest();
            if (object.Refid !== undefined && object.Refid !== null)
                message.Refid = object.Refid >>> 0;
            if (object.KmiID !== undefined && object.KmiID !== null)
                message.KmiID = object.KmiID >>> 0;
            return message;
        };

        /**
         * Creates a CreateDockerImageRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateDockerImageRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        CreateDockerImageRequest.from = CreateDockerImageRequest.fromObject;

        /**
         * Creates a plain object from a CreateDockerImageRequest message. Also converts values to other types if specified.
         * @param {customercontainer.CreateDockerImageRequest} message CreateDockerImageRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateDockerImageRequest.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.Refid = 0;
                object.KmiID = 0;
            }
            if (message.Refid !== undefined && message.Refid !== null && message.hasOwnProperty("Refid"))
                object.Refid = message.Refid;
            if (message.KmiID !== undefined && message.KmiID !== null && message.hasOwnProperty("KmiID"))
                object.KmiID = message.KmiID;
            return object;
        };

        /**
         * Creates a plain object from this CreateDockerImageRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateDockerImageRequest.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this CreateDockerImageRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        CreateDockerImageRequest.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CreateDockerImageRequest;
    })();

    customercontainer.CreateDockerImageResponse = (function() {

        /**
         * Constructs a new CreateDockerImageResponse.
         * @exports customercontainer.CreateDockerImageResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        function CreateDockerImageResponse(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    this[keys[i]] = properties[keys[i]];
        }

        /**
         * CreateDockerImageResponse ID.
         * @type {string|undefined}
         */
        CreateDockerImageResponse.prototype.ID = "";

        /**
         * CreateDockerImageResponse Error.
         * @type {string|undefined}
         */
        CreateDockerImageResponse.prototype.Error = "";

        /**
         * Creates a new CreateDockerImageResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse instance
         */
        CreateDockerImageResponse.create = function create(properties) {
            return new CreateDockerImageResponse(properties);
        };

        /**
         * Encodes the specified CreateDockerImageResponse message.
         * @param {customercontainer.CreateDockerImageResponse|Object} message CreateDockerImageResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateDockerImageResponse.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID !== undefined && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.ID);
            if (message.Error !== undefined && message.hasOwnProperty("Error"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.Error);
            return writer;
        };

        /**
         * Encodes the specified CreateDockerImageResponse message, length delimited.
         * @param {customercontainer.CreateDockerImageResponse|Object} message CreateDockerImageResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CreateDockerImageResponse.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CreateDockerImageResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        CreateDockerImageResponse.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.customercontainer.CreateDockerImageResponse();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.string();
                    break;
                case 2:
                    message.Error = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CreateDockerImageResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        CreateDockerImageResponse.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CreateDockerImageResponse message.
         * @param {customercontainer.CreateDockerImageResponse|Object} message CreateDockerImageResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        CreateDockerImageResponse.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID !== undefined)
                if (!$util.isString(message.ID))
                    return "ID: string expected";
            if (message.Error !== undefined)
                if (!$util.isString(message.Error))
                    return "Error: string expected";
            return null;
        };

        /**
         * Creates a CreateDockerImageResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        CreateDockerImageResponse.fromObject = function fromObject(object) {
            if (object instanceof $root.customercontainer.CreateDockerImageResponse)
                return object;
            let message = new $root.customercontainer.CreateDockerImageResponse();
            if (object.ID !== undefined && object.ID !== null)
                message.ID = String(object.ID);
            if (object.Error !== undefined && object.Error !== null)
                message.Error = String(object.Error);
            return message;
        };

        /**
         * Creates a CreateDockerImageResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateDockerImageResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        CreateDockerImageResponse.from = CreateDockerImageResponse.fromObject;

        /**
         * Creates a plain object from a CreateDockerImageResponse message. Also converts values to other types if specified.
         * @param {customercontainer.CreateDockerImageResponse} message CreateDockerImageResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateDockerImageResponse.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                object.ID = "";
                object.Error = "";
            }
            if (message.ID !== undefined && message.ID !== null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.Error !== undefined && message.Error !== null && message.hasOwnProperty("Error"))
                object.Error = message.Error;
            return object;
        };

        /**
         * Creates a plain object from this CreateDockerImageResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CreateDockerImageResponse.prototype.toObject = function toObject(options) {
            return this.constructor.toObject(this, options);
        };

        /**
         * Converts this CreateDockerImageResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        CreateDockerImageResponse.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CreateDockerImageResponse;
    })();

    return customercontainer;
})();

// Resolve lazy type references to actual types
$util.lazyResolve($root, $lazyTypes);

module.exports = $root;
