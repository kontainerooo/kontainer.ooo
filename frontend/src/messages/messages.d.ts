import * as $protobuf from "protobufjs";

/**
 * Namespace user.
 * @exports user
 * @namespace
 */
export namespace user {

    /**
     * Constructs a new UserService service.
     * @exports user.UserService
     * @extends $protobuf.rpc.Service
     * @constructor
     * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
     * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
     * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
     */
    class UserService extends $protobuf.rpc.Service {

        /**
         * Constructs a new UserService service.
         * @exports user.UserService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Creates new UserService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {UserService} RPC service. Useful where requests and/or responses are streamed.
         */
        static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): UserService;

        /**
         * Calls CreateUser.
         * @param {user.CreateUserRequest|Object} request CreateUserRequest message or plain object
         * @param {UserService_createUser_Callback} callback Node-style callback called with the error, if any, and CreateUserResponse
         * @returns {undefined}
         */
        createUser(request: (user.CreateUserRequest|Object), callback: UserService_createUser_Callback): void;

        /**
         * Calls EditUser.
         * @param {user.EditUserRequest|Object} request EditUserRequest message or plain object
         * @param {UserService_editUser_Callback} callback Node-style callback called with the error, if any, and EditUserResponse
         * @returns {undefined}
         */
        editUser(request: (user.EditUserRequest|Object), callback: UserService_editUser_Callback): void;

        /**
         * Calls ChangeUsername.
         * @param {user.ChangeUsernameRequest|Object} request ChangeUsernameRequest message or plain object
         * @param {UserService_changeUsername_Callback} callback Node-style callback called with the error, if any, and ChangeUsernameResponse
         * @returns {undefined}
         */
        changeUsername(request: (user.ChangeUsernameRequest|Object), callback: UserService_changeUsername_Callback): void;

        /**
         * Calls DeleteUser.
         * @param {user.DeleteUserRequest|Object} request DeleteUserRequest message or plain object
         * @param {UserService_deleteUser_Callback} callback Node-style callback called with the error, if any, and DeleteUserResponse
         * @returns {undefined}
         */
        deleteUser(request: (user.DeleteUserRequest|Object), callback: UserService_deleteUser_Callback): void;

        /**
         * Calls ResetPassword.
         * @param {user.ResetPasswordRequest|Object} request ResetPasswordRequest message or plain object
         * @param {UserService_resetPassword_Callback} callback Node-style callback called with the error, if any, and ResetPasswordResponse
         * @returns {undefined}
         */
        resetPassword(request: (user.ResetPasswordRequest|Object), callback: UserService_resetPassword_Callback): void;

        /**
         * Calls GetUser.
         * @param {user.GetUserRequest|Object} request GetUserRequest message or plain object
         * @param {UserService_getUser_Callback} callback Node-style callback called with the error, if any, and GetUserResponse
         * @returns {undefined}
         */
        getUser(request: (user.GetUserRequest|Object), callback: UserService_getUser_Callback): void;
    }

    /**
     * Constructs a new Address.
     * @exports user.Address
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class Address {

        /**
         * Constructs a new Address.
         * @exports user.Address
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * Address ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * Address postcode.
         * @type {string|undefined}
         */
        postcode?: string;

        /**
         * Address city.
         * @type {string|undefined}
         */
        city?: string;

        /**
         * Address country.
         * @type {string|undefined}
         */
        country?: string;

        /**
         * Address street.
         * @type {string|undefined}
         */
        street?: string;

        /**
         * Address houseno.
         * @type {number|undefined}
         */
        houseno?: number;

        /**
         * Address additional.
         * @type {string|undefined}
         */
        additional?: string;

        /**
         * Creates a new Address instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.Address} Address instance
         */
        static create(properties?: Object): user.Address;

        /**
         * Encodes the specified Address message.
         * @param {user.Address|Object} message Address message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.Address|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Address message, length delimited.
         * @param {user.Address|Object} message Address message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.Address|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an Address message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.Address} Address
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.Address;

        /**
         * Decodes an Address message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.Address} Address
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.Address;

        /**
         * Verifies an Address message.
         * @param {user.Address|Object} message Address message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.Address|Object)): string;

        /**
         * Creates an Address message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Address} Address
         */
        static fromObject(object: { [k: string]: any }): user.Address;

        /**
         * Creates an Address message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.Address.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Address} Address
         */
        static from(object: { [k: string]: any }): user.Address;

        /**
         * Creates a plain object from an Address message. Also converts values to other types if specified.
         * @param {user.Address} message Address
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.Address, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this Address message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this Address to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new Config.
     * @exports user.Config
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class Config {

        /**
         * Constructs a new Config.
         * @exports user.Config
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * Config admin.
         * @type {boolean|undefined}
         */
        admin?: boolean;

        /**
         * Config email.
         * @type {string|undefined}
         */
        email?: string;

        /**
         * Config password.
         * @type {string|undefined}
         */
        password?: string;

        /**
         * Config salt.
         * @type {string|undefined}
         */
        salt?: string;

        /**
         * Config Address.
         * @type {user.Address|undefined}
         */
        Address?: user.Address;

        /**
         * Config addressID.
         * @type {number|undefined}
         */
        addressID?: number;

        /**
         * Config phone.
         * @type {string|undefined}
         */
        phone?: string;

        /**
         * Config image.
         * @type {string|undefined}
         */
        image?: string;

        /**
         * Creates a new Config instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.Config} Config instance
         */
        static create(properties?: Object): user.Config;

        /**
         * Encodes the specified Config message.
         * @param {user.Config|Object} message Config message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.Config|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Config message, length delimited.
         * @param {user.Config|Object} message Config message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.Config|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Config message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.Config} Config
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.Config;

        /**
         * Decodes a Config message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.Config} Config
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.Config;

        /**
         * Verifies a Config message.
         * @param {user.Config|Object} message Config message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.Config|Object)): string;

        /**
         * Creates a Config message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Config} Config
         */
        static fromObject(object: { [k: string]: any }): user.Config;

        /**
         * Creates a Config message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.Config.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.Config} Config
         */
        static from(object: { [k: string]: any }): user.Config;

        /**
         * Creates a plain object from a Config message. Also converts values to other types if specified.
         * @param {user.Config} message Config
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.Config, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this Config message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this Config to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new User.
     * @exports user.User
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class User {

        /**
         * Constructs a new User.
         * @exports user.User
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * User ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * User username.
         * @type {string|undefined}
         */
        username?: string;

        /**
         * User config.
         * @type {user.Config|undefined}
         */
        config?: user.Config;

        /**
         * Creates a new User instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.User} User instance
         */
        static create(properties?: Object): user.User;

        /**
         * Encodes the specified User message.
         * @param {user.User|Object} message User message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.User|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified User message, length delimited.
         * @param {user.User|Object} message User message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.User|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a User message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.User} User
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.User;

        /**
         * Decodes a User message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.User} User
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.User;

        /**
         * Verifies a User message.
         * @param {user.User|Object} message User message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.User|Object)): string;

        /**
         * Creates a User message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.User} User
         */
        static fromObject(object: { [k: string]: any }): user.User;

        /**
         * Creates a User message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.User.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.User} User
         */
        static from(object: { [k: string]: any }): user.User;

        /**
         * Creates a plain object from a User message. Also converts values to other types if specified.
         * @param {user.User} message User
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.User, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this User message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this User to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new CreateUserRequest.
     * @exports user.CreateUserRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class CreateUserRequest {

        /**
         * Constructs a new CreateUserRequest.
         * @exports user.CreateUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * CreateUserRequest username.
         * @type {string|undefined}
         */
        username?: string;

        /**
         * CreateUserRequest config.
         * @type {user.Config|undefined}
         */
        config?: user.Config;

        /**
         * Creates a new CreateUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.CreateUserRequest} CreateUserRequest instance
         */
        static create(properties?: Object): user.CreateUserRequest;

        /**
         * Encodes the specified CreateUserRequest message.
         * @param {user.CreateUserRequest|Object} message CreateUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.CreateUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified CreateUserRequest message, length delimited.
         * @param {user.CreateUserRequest|Object} message CreateUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.CreateUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.CreateUserRequest;

        /**
         * Decodes a CreateUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.CreateUserRequest;

        /**
         * Verifies a CreateUserRequest message.
         * @param {user.CreateUserRequest|Object} message CreateUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.CreateUserRequest|Object)): string;

        /**
         * Creates a CreateUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        static fromObject(object: { [k: string]: any }): user.CreateUserRequest;

        /**
         * Creates a CreateUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.CreateUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserRequest} CreateUserRequest
         */
        static from(object: { [k: string]: any }): user.CreateUserRequest;

        /**
         * Creates a plain object from a CreateUserRequest message. Also converts values to other types if specified.
         * @param {user.CreateUserRequest} message CreateUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.CreateUserRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this CreateUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this CreateUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new CreateUserResponse.
     * @exports user.CreateUserResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class CreateUserResponse {

        /**
         * Constructs a new CreateUserResponse.
         * @exports user.CreateUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * CreateUserResponse ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * CreateUserResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new CreateUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.CreateUserResponse} CreateUserResponse instance
         */
        static create(properties?: Object): user.CreateUserResponse;

        /**
         * Encodes the specified CreateUserResponse message.
         * @param {user.CreateUserResponse|Object} message CreateUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.CreateUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified CreateUserResponse message, length delimited.
         * @param {user.CreateUserResponse|Object} message CreateUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.CreateUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.CreateUserResponse;

        /**
         * Decodes a CreateUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.CreateUserResponse;

        /**
         * Verifies a CreateUserResponse message.
         * @param {user.CreateUserResponse|Object} message CreateUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.CreateUserResponse|Object)): string;

        /**
         * Creates a CreateUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        static fromObject(object: { [k: string]: any }): user.CreateUserResponse;

        /**
         * Creates a CreateUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.CreateUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.CreateUserResponse} CreateUserResponse
         */
        static from(object: { [k: string]: any }): user.CreateUserResponse;

        /**
         * Creates a plain object from a CreateUserResponse message. Also converts values to other types if specified.
         * @param {user.CreateUserResponse} message CreateUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.CreateUserResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this CreateUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this CreateUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new EditUserRequest.
     * @exports user.EditUserRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class EditUserRequest {

        /**
         * Constructs a new EditUserRequest.
         * @exports user.EditUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * EditUserRequest ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * EditUserRequest config.
         * @type {user.Config|undefined}
         */
        config?: user.Config;

        /**
         * Creates a new EditUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.EditUserRequest} EditUserRequest instance
         */
        static create(properties?: Object): user.EditUserRequest;

        /**
         * Encodes the specified EditUserRequest message.
         * @param {user.EditUserRequest|Object} message EditUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.EditUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified EditUserRequest message, length delimited.
         * @param {user.EditUserRequest|Object} message EditUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.EditUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an EditUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.EditUserRequest} EditUserRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.EditUserRequest;

        /**
         * Decodes an EditUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.EditUserRequest} EditUserRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.EditUserRequest;

        /**
         * Verifies an EditUserRequest message.
         * @param {user.EditUserRequest|Object} message EditUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.EditUserRequest|Object)): string;

        /**
         * Creates an EditUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserRequest} EditUserRequest
         */
        static fromObject(object: { [k: string]: any }): user.EditUserRequest;

        /**
         * Creates an EditUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.EditUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserRequest} EditUserRequest
         */
        static from(object: { [k: string]: any }): user.EditUserRequest;

        /**
         * Creates a plain object from an EditUserRequest message. Also converts values to other types if specified.
         * @param {user.EditUserRequest} message EditUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.EditUserRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this EditUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this EditUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new EditUserResponse.
     * @exports user.EditUserResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class EditUserResponse {

        /**
         * Constructs a new EditUserResponse.
         * @exports user.EditUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * EditUserResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new EditUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.EditUserResponse} EditUserResponse instance
         */
        static create(properties?: Object): user.EditUserResponse;

        /**
         * Encodes the specified EditUserResponse message.
         * @param {user.EditUserResponse|Object} message EditUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.EditUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified EditUserResponse message, length delimited.
         * @param {user.EditUserResponse|Object} message EditUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.EditUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an EditUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.EditUserResponse} EditUserResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.EditUserResponse;

        /**
         * Decodes an EditUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.EditUserResponse} EditUserResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.EditUserResponse;

        /**
         * Verifies an EditUserResponse message.
         * @param {user.EditUserResponse|Object} message EditUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.EditUserResponse|Object)): string;

        /**
         * Creates an EditUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserResponse} EditUserResponse
         */
        static fromObject(object: { [k: string]: any }): user.EditUserResponse;

        /**
         * Creates an EditUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.EditUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.EditUserResponse} EditUserResponse
         */
        static from(object: { [k: string]: any }): user.EditUserResponse;

        /**
         * Creates a plain object from an EditUserResponse message. Also converts values to other types if specified.
         * @param {user.EditUserResponse} message EditUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.EditUserResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this EditUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this EditUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new ChangeUsernameRequest.
     * @exports user.ChangeUsernameRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class ChangeUsernameRequest {

        /**
         * Constructs a new ChangeUsernameRequest.
         * @exports user.ChangeUsernameRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * ChangeUsernameRequest ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * ChangeUsernameRequest username.
         * @type {string|undefined}
         */
        username?: string;

        /**
         * Creates a new ChangeUsernameRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest instance
         */
        static create(properties?: Object): user.ChangeUsernameRequest;

        /**
         * Encodes the specified ChangeUsernameRequest message.
         * @param {user.ChangeUsernameRequest|Object} message ChangeUsernameRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.ChangeUsernameRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ChangeUsernameRequest message, length delimited.
         * @param {user.ChangeUsernameRequest|Object} message ChangeUsernameRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.ChangeUsernameRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ChangeUsernameRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.ChangeUsernameRequest;

        /**
         * Decodes a ChangeUsernameRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.ChangeUsernameRequest;

        /**
         * Verifies a ChangeUsernameRequest message.
         * @param {user.ChangeUsernameRequest|Object} message ChangeUsernameRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.ChangeUsernameRequest|Object)): string;

        /**
         * Creates a ChangeUsernameRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        static fromObject(object: { [k: string]: any }): user.ChangeUsernameRequest;

        /**
         * Creates a ChangeUsernameRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ChangeUsernameRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameRequest} ChangeUsernameRequest
         */
        static from(object: { [k: string]: any }): user.ChangeUsernameRequest;

        /**
         * Creates a plain object from a ChangeUsernameRequest message. Also converts values to other types if specified.
         * @param {user.ChangeUsernameRequest} message ChangeUsernameRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.ChangeUsernameRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this ChangeUsernameRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this ChangeUsernameRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new ChangeUsernameResponse.
     * @exports user.ChangeUsernameResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class ChangeUsernameResponse {

        /**
         * Constructs a new ChangeUsernameResponse.
         * @exports user.ChangeUsernameResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * ChangeUsernameResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new ChangeUsernameResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse instance
         */
        static create(properties?: Object): user.ChangeUsernameResponse;

        /**
         * Encodes the specified ChangeUsernameResponse message.
         * @param {user.ChangeUsernameResponse|Object} message ChangeUsernameResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.ChangeUsernameResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ChangeUsernameResponse message, length delimited.
         * @param {user.ChangeUsernameResponse|Object} message ChangeUsernameResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.ChangeUsernameResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ChangeUsernameResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.ChangeUsernameResponse;

        /**
         * Decodes a ChangeUsernameResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.ChangeUsernameResponse;

        /**
         * Verifies a ChangeUsernameResponse message.
         * @param {user.ChangeUsernameResponse|Object} message ChangeUsernameResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.ChangeUsernameResponse|Object)): string;

        /**
         * Creates a ChangeUsernameResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        static fromObject(object: { [k: string]: any }): user.ChangeUsernameResponse;

        /**
         * Creates a ChangeUsernameResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ChangeUsernameResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ChangeUsernameResponse} ChangeUsernameResponse
         */
        static from(object: { [k: string]: any }): user.ChangeUsernameResponse;

        /**
         * Creates a plain object from a ChangeUsernameResponse message. Also converts values to other types if specified.
         * @param {user.ChangeUsernameResponse} message ChangeUsernameResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.ChangeUsernameResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this ChangeUsernameResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this ChangeUsernameResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new DeleteUserRequest.
     * @exports user.DeleteUserRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class DeleteUserRequest {

        /**
         * Constructs a new DeleteUserRequest.
         * @exports user.DeleteUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * DeleteUserRequest ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * Creates a new DeleteUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.DeleteUserRequest} DeleteUserRequest instance
         */
        static create(properties?: Object): user.DeleteUserRequest;

        /**
         * Encodes the specified DeleteUserRequest message.
         * @param {user.DeleteUserRequest|Object} message DeleteUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.DeleteUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified DeleteUserRequest message, length delimited.
         * @param {user.DeleteUserRequest|Object} message DeleteUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.DeleteUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.DeleteUserRequest;

        /**
         * Decodes a DeleteUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.DeleteUserRequest;

        /**
         * Verifies a DeleteUserRequest message.
         * @param {user.DeleteUserRequest|Object} message DeleteUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.DeleteUserRequest|Object)): string;

        /**
         * Creates a DeleteUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        static fromObject(object: { [k: string]: any }): user.DeleteUserRequest;

        /**
         * Creates a DeleteUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.DeleteUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserRequest} DeleteUserRequest
         */
        static from(object: { [k: string]: any }): user.DeleteUserRequest;

        /**
         * Creates a plain object from a DeleteUserRequest message. Also converts values to other types if specified.
         * @param {user.DeleteUserRequest} message DeleteUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.DeleteUserRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this DeleteUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this DeleteUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new DeleteUserResponse.
     * @exports user.DeleteUserResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class DeleteUserResponse {

        /**
         * Constructs a new DeleteUserResponse.
         * @exports user.DeleteUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * DeleteUserResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new DeleteUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.DeleteUserResponse} DeleteUserResponse instance
         */
        static create(properties?: Object): user.DeleteUserResponse;

        /**
         * Encodes the specified DeleteUserResponse message.
         * @param {user.DeleteUserResponse|Object} message DeleteUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.DeleteUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified DeleteUserResponse message, length delimited.
         * @param {user.DeleteUserResponse|Object} message DeleteUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.DeleteUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.DeleteUserResponse;

        /**
         * Decodes a DeleteUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.DeleteUserResponse;

        /**
         * Verifies a DeleteUserResponse message.
         * @param {user.DeleteUserResponse|Object} message DeleteUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.DeleteUserResponse|Object)): string;

        /**
         * Creates a DeleteUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        static fromObject(object: { [k: string]: any }): user.DeleteUserResponse;

        /**
         * Creates a DeleteUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.DeleteUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.DeleteUserResponse} DeleteUserResponse
         */
        static from(object: { [k: string]: any }): user.DeleteUserResponse;

        /**
         * Creates a plain object from a DeleteUserResponse message. Also converts values to other types if specified.
         * @param {user.DeleteUserResponse} message DeleteUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.DeleteUserResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this DeleteUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this DeleteUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new ResetPasswordRequest.
     * @exports user.ResetPasswordRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class ResetPasswordRequest {

        /**
         * Constructs a new ResetPasswordRequest.
         * @exports user.ResetPasswordRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * ResetPasswordRequest email.
         * @type {string|undefined}
         */
        email?: string;

        /**
         * Creates a new ResetPasswordRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest instance
         */
        static create(properties?: Object): user.ResetPasswordRequest;

        /**
         * Encodes the specified ResetPasswordRequest message.
         * @param {user.ResetPasswordRequest|Object} message ResetPasswordRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.ResetPasswordRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ResetPasswordRequest message, length delimited.
         * @param {user.ResetPasswordRequest|Object} message ResetPasswordRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.ResetPasswordRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ResetPasswordRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.ResetPasswordRequest;

        /**
         * Decodes a ResetPasswordRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.ResetPasswordRequest;

        /**
         * Verifies a ResetPasswordRequest message.
         * @param {user.ResetPasswordRequest|Object} message ResetPasswordRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.ResetPasswordRequest|Object)): string;

        /**
         * Creates a ResetPasswordRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        static fromObject(object: { [k: string]: any }): user.ResetPasswordRequest;

        /**
         * Creates a ResetPasswordRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ResetPasswordRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordRequest} ResetPasswordRequest
         */
        static from(object: { [k: string]: any }): user.ResetPasswordRequest;

        /**
         * Creates a plain object from a ResetPasswordRequest message. Also converts values to other types if specified.
         * @param {user.ResetPasswordRequest} message ResetPasswordRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.ResetPasswordRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this ResetPasswordRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this ResetPasswordRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new ResetPasswordResponse.
     * @exports user.ResetPasswordResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class ResetPasswordResponse {

        /**
         * Constructs a new ResetPasswordResponse.
         * @exports user.ResetPasswordResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * ResetPasswordResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new ResetPasswordResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse instance
         */
        static create(properties?: Object): user.ResetPasswordResponse;

        /**
         * Encodes the specified ResetPasswordResponse message.
         * @param {user.ResetPasswordResponse|Object} message ResetPasswordResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.ResetPasswordResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ResetPasswordResponse message, length delimited.
         * @param {user.ResetPasswordResponse|Object} message ResetPasswordResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.ResetPasswordResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ResetPasswordResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.ResetPasswordResponse;

        /**
         * Decodes a ResetPasswordResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.ResetPasswordResponse;

        /**
         * Verifies a ResetPasswordResponse message.
         * @param {user.ResetPasswordResponse|Object} message ResetPasswordResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.ResetPasswordResponse|Object)): string;

        /**
         * Creates a ResetPasswordResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        static fromObject(object: { [k: string]: any }): user.ResetPasswordResponse;

        /**
         * Creates a ResetPasswordResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.ResetPasswordResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.ResetPasswordResponse} ResetPasswordResponse
         */
        static from(object: { [k: string]: any }): user.ResetPasswordResponse;

        /**
         * Creates a plain object from a ResetPasswordResponse message. Also converts values to other types if specified.
         * @param {user.ResetPasswordResponse} message ResetPasswordResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.ResetPasswordResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this ResetPasswordResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this ResetPasswordResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new GetUserRequest.
     * @exports user.GetUserRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class GetUserRequest {

        /**
         * Constructs a new GetUserRequest.
         * @exports user.GetUserRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * GetUserRequest ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * Creates a new GetUserRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.GetUserRequest} GetUserRequest instance
         */
        static create(properties?: Object): user.GetUserRequest;

        /**
         * Encodes the specified GetUserRequest message.
         * @param {user.GetUserRequest|Object} message GetUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.GetUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified GetUserRequest message, length delimited.
         * @param {user.GetUserRequest|Object} message GetUserRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.GetUserRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GetUserRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.GetUserRequest} GetUserRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.GetUserRequest;

        /**
         * Decodes a GetUserRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.GetUserRequest} GetUserRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.GetUserRequest;

        /**
         * Verifies a GetUserRequest message.
         * @param {user.GetUserRequest|Object} message GetUserRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.GetUserRequest|Object)): string;

        /**
         * Creates a GetUserRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserRequest} GetUserRequest
         */
        static fromObject(object: { [k: string]: any }): user.GetUserRequest;

        /**
         * Creates a GetUserRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.GetUserRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserRequest} GetUserRequest
         */
        static from(object: { [k: string]: any }): user.GetUserRequest;

        /**
         * Creates a plain object from a GetUserRequest message. Also converts values to other types if specified.
         * @param {user.GetUserRequest} message GetUserRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.GetUserRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this GetUserRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this GetUserRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new GetUserResponse.
     * @exports user.GetUserResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class GetUserResponse {

        /**
         * Constructs a new GetUserResponse.
         * @exports user.GetUserResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * GetUserResponse user.
         * @type {user.User|undefined}
         */
        user?: user.User;

        /**
         * GetUserResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new GetUserResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {user.GetUserResponse} GetUserResponse instance
         */
        static create(properties?: Object): user.GetUserResponse;

        /**
         * Encodes the specified GetUserResponse message.
         * @param {user.GetUserResponse|Object} message GetUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (user.GetUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified GetUserResponse message, length delimited.
         * @param {user.GetUserResponse|Object} message GetUserResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (user.GetUserResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GetUserResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {user.GetUserResponse} GetUserResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): user.GetUserResponse;

        /**
         * Decodes a GetUserResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {user.GetUserResponse} GetUserResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): user.GetUserResponse;

        /**
         * Verifies a GetUserResponse message.
         * @param {user.GetUserResponse|Object} message GetUserResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (user.GetUserResponse|Object)): string;

        /**
         * Creates a GetUserResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserResponse} GetUserResponse
         */
        static fromObject(object: { [k: string]: any }): user.GetUserResponse;

        /**
         * Creates a GetUserResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link user.GetUserResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {user.GetUserResponse} GetUserResponse
         */
        static from(object: { [k: string]: any }): user.GetUserResponse;

        /**
         * Creates a plain object from a GetUserResponse message. Also converts values to other types if specified.
         * @param {user.GetUserResponse} message GetUserResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: user.GetUserResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this GetUserResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this GetUserResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }
}

/**
 * Callback as used by {@link UserService#createUser}.
 * @typedef UserService_createUser_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {user.CreateUserResponse} [response] CreateUserResponse
 */
type UserService_createUser_Callback = (error: Error, response?: user.CreateUserResponse) => void;

/**
 * Callback as used by {@link UserService#editUser}.
 * @typedef UserService_editUser_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {user.EditUserResponse} [response] EditUserResponse
 */
type UserService_editUser_Callback = (error: Error, response?: user.EditUserResponse) => void;

/**
 * Callback as used by {@link UserService#changeUsername}.
 * @typedef UserService_changeUsername_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {user.ChangeUsernameResponse} [response] ChangeUsernameResponse
 */
type UserService_changeUsername_Callback = (error: Error, response?: user.ChangeUsernameResponse) => void;

/**
 * Callback as used by {@link UserService#deleteUser}.
 * @typedef UserService_deleteUser_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {user.DeleteUserResponse} [response] DeleteUserResponse
 */
type UserService_deleteUser_Callback = (error: Error, response?: user.DeleteUserResponse) => void;

/**
 * Callback as used by {@link UserService#resetPassword}.
 * @typedef UserService_resetPassword_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {user.ResetPasswordResponse} [response] ResetPasswordResponse
 */
type UserService_resetPassword_Callback = (error: Error, response?: user.ResetPasswordResponse) => void;

/**
 * Callback as used by {@link UserService#getUser}.
 * @typedef UserService_getUser_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {user.GetUserResponse} [response] GetUserResponse
 */
type UserService_getUser_Callback = (error: Error, response?: user.GetUserResponse) => void;

/**
 * Namespace kmi.
 * @exports kmi
 * @namespace
 */
export namespace kmi {

    /**
     * Constructs a new KMIService service.
     * @exports kmi.KMIService
     * @extends $protobuf.rpc.Service
     * @constructor
     * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
     * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
     * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
     */
    class KMIService extends $protobuf.rpc.Service {

        /**
         * Constructs a new KMIService service.
         * @exports kmi.KMIService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Creates new KMIService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {KMIService} RPC service. Useful where requests and/or responses are streamed.
         */
        static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): KMIService;

        /**
         * Calls AddKMI.
         * @param {kmi.AddKMIRequest|Object} request AddKMIRequest message or plain object
         * @param {KMIService_addKMI_Callback} callback Node-style callback called with the error, if any, and AddKMIResponse
         * @returns {undefined}
         */
        addKMI(request: (kmi.AddKMIRequest|Object), callback: KMIService_addKMI_Callback): void;

        /**
         * Calls RemoveKMI.
         * @param {kmi.RemoveKMIRequest|Object} request RemoveKMIRequest message or plain object
         * @param {KMIService_removeKMI_Callback} callback Node-style callback called with the error, if any, and RemoveKMIResponse
         * @returns {undefined}
         */
        removeKMI(request: (kmi.RemoveKMIRequest|Object), callback: KMIService_removeKMI_Callback): void;

        /**
         * Calls GetKMI.
         * @param {kmi.GetKMIRequest|Object} request GetKMIRequest message or plain object
         * @param {KMIService_getKMI_Callback} callback Node-style callback called with the error, if any, and GetKMIResponse
         * @returns {undefined}
         */
        getKMI(request: (kmi.GetKMIRequest|Object), callback: KMIService_getKMI_Callback): void;

        /**
         * Calls KMI.
         * @param {kmi.KMIRequest|Object} request KMIRequest message or plain object
         * @param {KMIService_kMI_Callback} callback Node-style callback called with the error, if any, and KMIResponse
         * @returns {undefined}
         */
        kMI(request: (kmi.KMIRequest|Object), callback: KMIService_kMI_Callback): void;
    }

    /**
     * Type enum.
     * @name Type
     * @memberof kmi
     * @enum {number}
     * @property {number} WEBSERVER=0 WEBSERVER value
     */
    enum Type {
        WEBSERVER = 0
    }

    /**
     * Constructs a new KMDI.
     * @exports kmi.KMDI
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class KMDI {

        /**
         * Constructs a new KMDI.
         * @exports kmi.KMDI
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * KMDI ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * KMDI name.
         * @type {string|undefined}
         */
        name?: string;

        /**
         * KMDI version.
         * @type {string|undefined}
         */
        version?: string;

        /**
         * KMDI description.
         * @type {string|undefined}
         */
        description?: string;

        /**
         * KMDI type.
         * @type {number|undefined}
         */
        type?: number;

        /**
         * Creates a new KMDI instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMDI} KMDI instance
         */
        static create(properties?: Object): kmi.KMDI;

        /**
         * Encodes the specified KMDI message.
         * @param {kmi.KMDI|Object} message KMDI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.KMDI|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified KMDI message, length delimited.
         * @param {kmi.KMDI|Object} message KMDI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.KMDI|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a KMDI message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMDI} KMDI
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.KMDI;

        /**
         * Decodes a KMDI message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMDI} KMDI
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.KMDI;

        /**
         * Verifies a KMDI message.
         * @param {kmi.KMDI|Object} message KMDI message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.KMDI|Object)): string;

        /**
         * Creates a KMDI message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMDI} KMDI
         */
        static fromObject(object: { [k: string]: any }): kmi.KMDI;

        /**
         * Creates a KMDI message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMDI.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMDI} KMDI
         */
        static from(object: { [k: string]: any }): kmi.KMDI;

        /**
         * Creates a plain object from a KMDI message. Also converts values to other types if specified.
         * @param {kmi.KMDI} message KMDI
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.KMDI, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this KMDI message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this KMDI to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new FrontendModule.
     * @exports kmi.FrontendModule
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class FrontendModule {

        /**
         * Constructs a new FrontendModule.
         * @exports kmi.FrontendModule
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * FrontendModule template.
         * @type {string|undefined}
         */
        template?: string;

        /**
         * FrontendModule parameters.
         * @type {Object.<string,string>|undefined}
         */
        parameters?: { [k: string]: string };

        /**
         * Creates a new FrontendModule instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.FrontendModule} FrontendModule instance
         */
        static create(properties?: Object): kmi.FrontendModule;

        /**
         * Encodes the specified FrontendModule message.
         * @param {kmi.FrontendModule|Object} message FrontendModule message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.FrontendModule|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified FrontendModule message, length delimited.
         * @param {kmi.FrontendModule|Object} message FrontendModule message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.FrontendModule|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a FrontendModule message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.FrontendModule} FrontendModule
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.FrontendModule;

        /**
         * Decodes a FrontendModule message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.FrontendModule} FrontendModule
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.FrontendModule;

        /**
         * Verifies a FrontendModule message.
         * @param {kmi.FrontendModule|Object} message FrontendModule message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.FrontendModule|Object)): string;

        /**
         * Creates a FrontendModule message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.FrontendModule} FrontendModule
         */
        static fromObject(object: { [k: string]: any }): kmi.FrontendModule;

        /**
         * Creates a FrontendModule message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.FrontendModule.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.FrontendModule} FrontendModule
         */
        static from(object: { [k: string]: any }): kmi.FrontendModule;

        /**
         * Creates a plain object from a FrontendModule message. Also converts values to other types if specified.
         * @param {kmi.FrontendModule} message FrontendModule
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.FrontendModule, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this FrontendModule message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this FrontendModule to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new KMI.
     * @exports kmi.KMI
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class KMI {

        /**
         * Constructs a new KMI.
         * @exports kmi.KMI
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * KMI KMDI.
         * @type {kmi.KMDI|undefined}
         */
        KMDI?: kmi.KMDI;

        /**
         * KMI dockerfile.
         * @type {string|undefined}
         */
        dockerfile?: string;

        /**
         * KMI container.
         * @type {string|undefined}
         */
        container?: string;

        /**
         * KMI commands.
         * @type {Object.<string,string>|undefined}
         */
        commands?: { [k: string]: string };

        /**
         * KMI environment.
         * @type {Object.<string,string>|undefined}
         */
        environment?: { [k: string]: string };

        /**
         * KMI frontend.
         * @type {Array.<kmi.FrontendModule>|undefined}
         */
        frontend?: kmi.FrontendModule[];

        /**
         * KMI imports.
         * @type {Array.<string>|undefined}
         */
        imports?: string[];

        /**
         * KMI interfaces.
         * @type {Object.<string,string>|undefined}
         */
        interfaces?: { [k: string]: string };

        /**
         * KMI mounts.
         * @type {Array.<string>|undefined}
         */
        mounts?: string[];

        /**
         * KMI variables.
         * @type {Array.<string>|undefined}
         */
        variables?: string[];

        /**
         * KMI resources.
         * @type {Object.<string,string>|undefined}
         */
        resources?: { [k: string]: string };

        /**
         * Creates a new KMI instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMI} KMI instance
         */
        static create(properties?: Object): kmi.KMI;

        /**
         * Encodes the specified KMI message.
         * @param {kmi.KMI|Object} message KMI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.KMI|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified KMI message, length delimited.
         * @param {kmi.KMI|Object} message KMI message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.KMI|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a KMI message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMI} KMI
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.KMI;

        /**
         * Decodes a KMI message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMI} KMI
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.KMI;

        /**
         * Verifies a KMI message.
         * @param {kmi.KMI|Object} message KMI message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.KMI|Object)): string;

        /**
         * Creates a KMI message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMI} KMI
         */
        static fromObject(object: { [k: string]: any }): kmi.KMI;

        /**
         * Creates a KMI message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMI.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMI} KMI
         */
        static from(object: { [k: string]: any }): kmi.KMI;

        /**
         * Creates a plain object from a KMI message. Also converts values to other types if specified.
         * @param {kmi.KMI} message KMI
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.KMI, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this KMI message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this KMI to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new AddKMIRequest.
     * @exports kmi.AddKMIRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class AddKMIRequest {

        /**
         * Constructs a new AddKMIRequest.
         * @exports kmi.AddKMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * AddKMIRequest path.
         * @type {string|undefined}
         */
        path?: string;

        /**
         * Creates a new AddKMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.AddKMIRequest} AddKMIRequest instance
         */
        static create(properties?: Object): kmi.AddKMIRequest;

        /**
         * Encodes the specified AddKMIRequest message.
         * @param {kmi.AddKMIRequest|Object} message AddKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.AddKMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified AddKMIRequest message, length delimited.
         * @param {kmi.AddKMIRequest|Object} message AddKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.AddKMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AddKMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.AddKMIRequest;

        /**
         * Decodes an AddKMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.AddKMIRequest;

        /**
         * Verifies an AddKMIRequest message.
         * @param {kmi.AddKMIRequest|Object} message AddKMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.AddKMIRequest|Object)): string;

        /**
         * Creates an AddKMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        static fromObject(object: { [k: string]: any }): kmi.AddKMIRequest;

        /**
         * Creates an AddKMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.AddKMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIRequest} AddKMIRequest
         */
        static from(object: { [k: string]: any }): kmi.AddKMIRequest;

        /**
         * Creates a plain object from an AddKMIRequest message. Also converts values to other types if specified.
         * @param {kmi.AddKMIRequest} message AddKMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.AddKMIRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this AddKMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this AddKMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new AddKMIResponse.
     * @exports kmi.AddKMIResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class AddKMIResponse {

        /**
         * Constructs a new AddKMIResponse.
         * @exports kmi.AddKMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * AddKMIResponse ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * AddKMIResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new AddKMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.AddKMIResponse} AddKMIResponse instance
         */
        static create(properties?: Object): kmi.AddKMIResponse;

        /**
         * Encodes the specified AddKMIResponse message.
         * @param {kmi.AddKMIResponse|Object} message AddKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.AddKMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified AddKMIResponse message, length delimited.
         * @param {kmi.AddKMIResponse|Object} message AddKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.AddKMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AddKMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.AddKMIResponse;

        /**
         * Decodes an AddKMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.AddKMIResponse;

        /**
         * Verifies an AddKMIResponse message.
         * @param {kmi.AddKMIResponse|Object} message AddKMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.AddKMIResponse|Object)): string;

        /**
         * Creates an AddKMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        static fromObject(object: { [k: string]: any }): kmi.AddKMIResponse;

        /**
         * Creates an AddKMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.AddKMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.AddKMIResponse} AddKMIResponse
         */
        static from(object: { [k: string]: any }): kmi.AddKMIResponse;

        /**
         * Creates a plain object from an AddKMIResponse message. Also converts values to other types if specified.
         * @param {kmi.AddKMIResponse} message AddKMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.AddKMIResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this AddKMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this AddKMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new RemoveKMIRequest.
     * @exports kmi.RemoveKMIRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class RemoveKMIRequest {

        /**
         * Constructs a new RemoveKMIRequest.
         * @exports kmi.RemoveKMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * RemoveKMIRequest ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * Creates a new RemoveKMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest instance
         */
        static create(properties?: Object): kmi.RemoveKMIRequest;

        /**
         * Encodes the specified RemoveKMIRequest message.
         * @param {kmi.RemoveKMIRequest|Object} message RemoveKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.RemoveKMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified RemoveKMIRequest message, length delimited.
         * @param {kmi.RemoveKMIRequest|Object} message RemoveKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.RemoveKMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a RemoveKMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.RemoveKMIRequest;

        /**
         * Decodes a RemoveKMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.RemoveKMIRequest;

        /**
         * Verifies a RemoveKMIRequest message.
         * @param {kmi.RemoveKMIRequest|Object} message RemoveKMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.RemoveKMIRequest|Object)): string;

        /**
         * Creates a RemoveKMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        static fromObject(object: { [k: string]: any }): kmi.RemoveKMIRequest;

        /**
         * Creates a RemoveKMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.RemoveKMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIRequest} RemoveKMIRequest
         */
        static from(object: { [k: string]: any }): kmi.RemoveKMIRequest;

        /**
         * Creates a plain object from a RemoveKMIRequest message. Also converts values to other types if specified.
         * @param {kmi.RemoveKMIRequest} message RemoveKMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.RemoveKMIRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this RemoveKMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this RemoveKMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new RemoveKMIResponse.
     * @exports kmi.RemoveKMIResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class RemoveKMIResponse {

        /**
         * Constructs a new RemoveKMIResponse.
         * @exports kmi.RemoveKMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * RemoveKMIResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new RemoveKMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse instance
         */
        static create(properties?: Object): kmi.RemoveKMIResponse;

        /**
         * Encodes the specified RemoveKMIResponse message.
         * @param {kmi.RemoveKMIResponse|Object} message RemoveKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.RemoveKMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified RemoveKMIResponse message, length delimited.
         * @param {kmi.RemoveKMIResponse|Object} message RemoveKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.RemoveKMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a RemoveKMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.RemoveKMIResponse;

        /**
         * Decodes a RemoveKMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.RemoveKMIResponse;

        /**
         * Verifies a RemoveKMIResponse message.
         * @param {kmi.RemoveKMIResponse|Object} message RemoveKMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.RemoveKMIResponse|Object)): string;

        /**
         * Creates a RemoveKMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        static fromObject(object: { [k: string]: any }): kmi.RemoveKMIResponse;

        /**
         * Creates a RemoveKMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.RemoveKMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.RemoveKMIResponse} RemoveKMIResponse
         */
        static from(object: { [k: string]: any }): kmi.RemoveKMIResponse;

        /**
         * Creates a plain object from a RemoveKMIResponse message. Also converts values to other types if specified.
         * @param {kmi.RemoveKMIResponse} message RemoveKMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.RemoveKMIResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this RemoveKMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this RemoveKMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new GetKMIRequest.
     * @exports kmi.GetKMIRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class GetKMIRequest {

        /**
         * Constructs a new GetKMIRequest.
         * @exports kmi.GetKMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * GetKMIRequest ID.
         * @type {number|undefined}
         */
        ID?: number;

        /**
         * Creates a new GetKMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.GetKMIRequest} GetKMIRequest instance
         */
        static create(properties?: Object): kmi.GetKMIRequest;

        /**
         * Encodes the specified GetKMIRequest message.
         * @param {kmi.GetKMIRequest|Object} message GetKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.GetKMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified GetKMIRequest message, length delimited.
         * @param {kmi.GetKMIRequest|Object} message GetKMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.GetKMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GetKMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.GetKMIRequest;

        /**
         * Decodes a GetKMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.GetKMIRequest;

        /**
         * Verifies a GetKMIRequest message.
         * @param {kmi.GetKMIRequest|Object} message GetKMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.GetKMIRequest|Object)): string;

        /**
         * Creates a GetKMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        static fromObject(object: { [k: string]: any }): kmi.GetKMIRequest;

        /**
         * Creates a GetKMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.GetKMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIRequest} GetKMIRequest
         */
        static from(object: { [k: string]: any }): kmi.GetKMIRequest;

        /**
         * Creates a plain object from a GetKMIRequest message. Also converts values to other types if specified.
         * @param {kmi.GetKMIRequest} message GetKMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.GetKMIRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this GetKMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this GetKMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new GetKMIResponse.
     * @exports kmi.GetKMIResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class GetKMIResponse {

        /**
         * Constructs a new GetKMIResponse.
         * @exports kmi.GetKMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * GetKMIResponse kmi.
         * @type {kmi.KMI|undefined}
         */
        kmi?: kmi.KMI;

        /**
         * GetKMIResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new GetKMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.GetKMIResponse} GetKMIResponse instance
         */
        static create(properties?: Object): kmi.GetKMIResponse;

        /**
         * Encodes the specified GetKMIResponse message.
         * @param {kmi.GetKMIResponse|Object} message GetKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.GetKMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified GetKMIResponse message, length delimited.
         * @param {kmi.GetKMIResponse|Object} message GetKMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.GetKMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GetKMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.GetKMIResponse;

        /**
         * Decodes a GetKMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.GetKMIResponse;

        /**
         * Verifies a GetKMIResponse message.
         * @param {kmi.GetKMIResponse|Object} message GetKMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.GetKMIResponse|Object)): string;

        /**
         * Creates a GetKMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        static fromObject(object: { [k: string]: any }): kmi.GetKMIResponse;

        /**
         * Creates a GetKMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.GetKMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.GetKMIResponse} GetKMIResponse
         */
        static from(object: { [k: string]: any }): kmi.GetKMIResponse;

        /**
         * Creates a plain object from a GetKMIResponse message. Also converts values to other types if specified.
         * @param {kmi.GetKMIResponse} message GetKMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.GetKMIResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this GetKMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this GetKMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new KMIRequest.
     * @exports kmi.KMIRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class KMIRequest {

        /**
         * Constructs a new KMIRequest.
         * @exports kmi.KMIRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * Creates a new KMIRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMIRequest} KMIRequest instance
         */
        static create(properties?: Object): kmi.KMIRequest;

        /**
         * Encodes the specified KMIRequest message.
         * @param {kmi.KMIRequest|Object} message KMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.KMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified KMIRequest message, length delimited.
         * @param {kmi.KMIRequest|Object} message KMIRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.KMIRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a KMIRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMIRequest} KMIRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.KMIRequest;

        /**
         * Decodes a KMIRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMIRequest} KMIRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.KMIRequest;

        /**
         * Verifies a KMIRequest message.
         * @param {kmi.KMIRequest|Object} message KMIRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.KMIRequest|Object)): string;

        /**
         * Creates a KMIRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIRequest} KMIRequest
         */
        static fromObject(object: { [k: string]: any }): kmi.KMIRequest;

        /**
         * Creates a KMIRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMIRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIRequest} KMIRequest
         */
        static from(object: { [k: string]: any }): kmi.KMIRequest;

        /**
         * Creates a plain object from a KMIRequest message. Also converts values to other types if specified.
         * @param {kmi.KMIRequest} message KMIRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.KMIRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this KMIRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this KMIRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new KMIResponse.
     * @exports kmi.KMIResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class KMIResponse {

        /**
         * Constructs a new KMIResponse.
         * @exports kmi.KMIResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * KMIResponse kmdi.
         * @type {Array.<kmi.KMDI>|undefined}
         */
        kmdi?: kmi.KMDI[];

        /**
         * KMIResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new KMIResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {kmi.KMIResponse} KMIResponse instance
         */
        static create(properties?: Object): kmi.KMIResponse;

        /**
         * Encodes the specified KMIResponse message.
         * @param {kmi.KMIResponse|Object} message KMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (kmi.KMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified KMIResponse message, length delimited.
         * @param {kmi.KMIResponse|Object} message KMIResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (kmi.KMIResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a KMIResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {kmi.KMIResponse} KMIResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): kmi.KMIResponse;

        /**
         * Decodes a KMIResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {kmi.KMIResponse} KMIResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): kmi.KMIResponse;

        /**
         * Verifies a KMIResponse message.
         * @param {kmi.KMIResponse|Object} message KMIResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (kmi.KMIResponse|Object)): string;

        /**
         * Creates a KMIResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIResponse} KMIResponse
         */
        static fromObject(object: { [k: string]: any }): kmi.KMIResponse;

        /**
         * Creates a KMIResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link kmi.KMIResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {kmi.KMIResponse} KMIResponse
         */
        static from(object: { [k: string]: any }): kmi.KMIResponse;

        /**
         * Creates a plain object from a KMIResponse message. Also converts values to other types if specified.
         * @param {kmi.KMIResponse} message KMIResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: kmi.KMIResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this KMIResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this KMIResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }
}

/**
 * Callback as used by {@link KMIService#addKMI}.
 * @typedef KMIService_addKMI_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {kmi.AddKMIResponse} [response] AddKMIResponse
 */
type KMIService_addKMI_Callback = (error: Error, response?: kmi.AddKMIResponse) => void;

/**
 * Callback as used by {@link KMIService#removeKMI}.
 * @typedef KMIService_removeKMI_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {kmi.RemoveKMIResponse} [response] RemoveKMIResponse
 */
type KMIService_removeKMI_Callback = (error: Error, response?: kmi.RemoveKMIResponse) => void;

/**
 * Callback as used by {@link KMIService#getKMI}.
 * @typedef KMIService_getKMI_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {kmi.GetKMIResponse} [response] GetKMIResponse
 */
type KMIService_getKMI_Callback = (error: Error, response?: kmi.GetKMIResponse) => void;

/**
 * Callback as used by {@link KMIService#kMI}.
 * @typedef KMIService_kMI_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {kmi.KMIResponse} [response] KMIResponse
 */
type KMIService_kMI_Callback = (error: Error, response?: kmi.KMIResponse) => void;

/**
 * Namespace containerlifecycle.
 * @exports containerlifecycle
 * @namespace
 */
export namespace containerlifecycle {

    /**
     * Constructs a new ContainerLifecycleService service.
     * @exports containerlifecycle.ContainerLifecycleService
     * @extends $protobuf.rpc.Service
     * @constructor
     * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
     * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
     * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
     */
    class ContainerLifecycleService extends $protobuf.rpc.Service {

        /**
         * Constructs a new ContainerLifecycleService service.
         * @exports containerlifecycle.ContainerLifecycleService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Creates new ContainerLifecycleService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {ContainerLifecycleService} RPC service. Useful where requests and/or responses are streamed.
         */
        static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): ContainerLifecycleService;

        /**
         * Calls StartContainer.
         * @param {containerlifecycle.StartContainerRequest|Object} request StartContainerRequest message or plain object
         * @param {ContainerLifecycleService_startContainer_Callback} callback Node-style callback called with the error, if any, and StartContainerResponse
         * @returns {undefined}
         */
        startContainer(request: (containerlifecycle.StartContainerRequest|Object), callback: ContainerLifecycleService_startContainer_Callback): void;

        /**
         * Calls StartCommand.
         * @param {containerlifecycle.StartCommandRequest|Object} request StartCommandRequest message or plain object
         * @param {ContainerLifecycleService_startCommand_Callback} callback Node-style callback called with the error, if any, and StartCommandResponse
         * @returns {undefined}
         */
        startCommand(request: (containerlifecycle.StartCommandRequest|Object), callback: ContainerLifecycleService_startCommand_Callback): void;

        /**
         * Calls StopContainer.
         * @param {containerlifecycle.StopContainerRequest|Object} request StopContainerRequest message or plain object
         * @param {ContainerLifecycleService_stopContainer_Callback} callback Node-style callback called with the error, if any, and StopContainerResponse
         * @returns {undefined}
         */
        stopContainer(request: (containerlifecycle.StopContainerRequest|Object), callback: ContainerLifecycleService_stopContainer_Callback): void;
    }

    /**
     * Constructs a new StartContainerRequest.
     * @exports containerlifecycle.StartContainerRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class StartContainerRequest {

        /**
         * Constructs a new StartContainerRequest.
         * @exports containerlifecycle.StartContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * StartContainerRequest ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * Creates a new StartContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest instance
         */
        static create(properties?: Object): containerlifecycle.StartContainerRequest;

        /**
         * Encodes the specified StartContainerRequest message.
         * @param {containerlifecycle.StartContainerRequest|Object} message StartContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (containerlifecycle.StartContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified StartContainerRequest message, length delimited.
         * @param {containerlifecycle.StartContainerRequest|Object} message StartContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (containerlifecycle.StartContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StartContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): containerlifecycle.StartContainerRequest;

        /**
         * Decodes a StartContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): containerlifecycle.StartContainerRequest;

        /**
         * Verifies a StartContainerRequest message.
         * @param {containerlifecycle.StartContainerRequest|Object} message StartContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (containerlifecycle.StartContainerRequest|Object)): string;

        /**
         * Creates a StartContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        static fromObject(object: { [k: string]: any }): containerlifecycle.StartContainerRequest;

        /**
         * Creates a StartContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerRequest} StartContainerRequest
         */
        static from(object: { [k: string]: any }): containerlifecycle.StartContainerRequest;

        /**
         * Creates a plain object from a StartContainerRequest message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartContainerRequest} message StartContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: containerlifecycle.StartContainerRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this StartContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this StartContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new StartCommandRequest.
     * @exports containerlifecycle.StartCommandRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class StartCommandRequest {

        /**
         * Constructs a new StartCommandRequest.
         * @exports containerlifecycle.StartCommandRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * StartCommandRequest ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * StartCommandRequest cmd.
         * @type {string|undefined}
         */
        cmd?: string;

        /**
         * Creates a new StartCommandRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest instance
         */
        static create(properties?: Object): containerlifecycle.StartCommandRequest;

        /**
         * Encodes the specified StartCommandRequest message.
         * @param {containerlifecycle.StartCommandRequest|Object} message StartCommandRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (containerlifecycle.StartCommandRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified StartCommandRequest message, length delimited.
         * @param {containerlifecycle.StartCommandRequest|Object} message StartCommandRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (containerlifecycle.StartCommandRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StartCommandRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): containerlifecycle.StartCommandRequest;

        /**
         * Decodes a StartCommandRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): containerlifecycle.StartCommandRequest;

        /**
         * Verifies a StartCommandRequest message.
         * @param {containerlifecycle.StartCommandRequest|Object} message StartCommandRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (containerlifecycle.StartCommandRequest|Object)): string;

        /**
         * Creates a StartCommandRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        static fromObject(object: { [k: string]: any }): containerlifecycle.StartCommandRequest;

        /**
         * Creates a StartCommandRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartCommandRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandRequest} StartCommandRequest
         */
        static from(object: { [k: string]: any }): containerlifecycle.StartCommandRequest;

        /**
         * Creates a plain object from a StartCommandRequest message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartCommandRequest} message StartCommandRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: containerlifecycle.StartCommandRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this StartCommandRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this StartCommandRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new StopContainerRequest.
     * @exports containerlifecycle.StopContainerRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class StopContainerRequest {

        /**
         * Constructs a new StopContainerRequest.
         * @exports containerlifecycle.StopContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * StopContainerRequest ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * Creates a new StopContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest instance
         */
        static create(properties?: Object): containerlifecycle.StopContainerRequest;

        /**
         * Encodes the specified StopContainerRequest message.
         * @param {containerlifecycle.StopContainerRequest|Object} message StopContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (containerlifecycle.StopContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified StopContainerRequest message, length delimited.
         * @param {containerlifecycle.StopContainerRequest|Object} message StopContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (containerlifecycle.StopContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StopContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): containerlifecycle.StopContainerRequest;

        /**
         * Decodes a StopContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): containerlifecycle.StopContainerRequest;

        /**
         * Verifies a StopContainerRequest message.
         * @param {containerlifecycle.StopContainerRequest|Object} message StopContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (containerlifecycle.StopContainerRequest|Object)): string;

        /**
         * Creates a StopContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        static fromObject(object: { [k: string]: any }): containerlifecycle.StopContainerRequest;

        /**
         * Creates a StopContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StopContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerRequest} StopContainerRequest
         */
        static from(object: { [k: string]: any }): containerlifecycle.StopContainerRequest;

        /**
         * Creates a plain object from a StopContainerRequest message. Also converts values to other types if specified.
         * @param {containerlifecycle.StopContainerRequest} message StopContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: containerlifecycle.StopContainerRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this StopContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this StopContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new StartContainerResponse.
     * @exports containerlifecycle.StartContainerResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class StartContainerResponse {

        /**
         * Constructs a new StartContainerResponse.
         * @exports containerlifecycle.StartContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * StartContainerResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new StartContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse instance
         */
        static create(properties?: Object): containerlifecycle.StartContainerResponse;

        /**
         * Encodes the specified StartContainerResponse message.
         * @param {containerlifecycle.StartContainerResponse|Object} message StartContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (containerlifecycle.StartContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified StartContainerResponse message, length delimited.
         * @param {containerlifecycle.StartContainerResponse|Object} message StartContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (containerlifecycle.StartContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StartContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): containerlifecycle.StartContainerResponse;

        /**
         * Decodes a StartContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): containerlifecycle.StartContainerResponse;

        /**
         * Verifies a StartContainerResponse message.
         * @param {containerlifecycle.StartContainerResponse|Object} message StartContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (containerlifecycle.StartContainerResponse|Object)): string;

        /**
         * Creates a StartContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        static fromObject(object: { [k: string]: any }): containerlifecycle.StartContainerResponse;

        /**
         * Creates a StartContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartContainerResponse} StartContainerResponse
         */
        static from(object: { [k: string]: any }): containerlifecycle.StartContainerResponse;

        /**
         * Creates a plain object from a StartContainerResponse message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartContainerResponse} message StartContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: containerlifecycle.StartContainerResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this StartContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this StartContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new StartCommandResponse.
     * @exports containerlifecycle.StartCommandResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class StartCommandResponse {

        /**
         * Constructs a new StartCommandResponse.
         * @exports containerlifecycle.StartCommandResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * StartCommandResponse ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * StartCommandResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new StartCommandResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse instance
         */
        static create(properties?: Object): containerlifecycle.StartCommandResponse;

        /**
         * Encodes the specified StartCommandResponse message.
         * @param {containerlifecycle.StartCommandResponse|Object} message StartCommandResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (containerlifecycle.StartCommandResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified StartCommandResponse message, length delimited.
         * @param {containerlifecycle.StartCommandResponse|Object} message StartCommandResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (containerlifecycle.StartCommandResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StartCommandResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): containerlifecycle.StartCommandResponse;

        /**
         * Decodes a StartCommandResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): containerlifecycle.StartCommandResponse;

        /**
         * Verifies a StartCommandResponse message.
         * @param {containerlifecycle.StartCommandResponse|Object} message StartCommandResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (containerlifecycle.StartCommandResponse|Object)): string;

        /**
         * Creates a StartCommandResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        static fromObject(object: { [k: string]: any }): containerlifecycle.StartCommandResponse;

        /**
         * Creates a StartCommandResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StartCommandResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StartCommandResponse} StartCommandResponse
         */
        static from(object: { [k: string]: any }): containerlifecycle.StartCommandResponse;

        /**
         * Creates a plain object from a StartCommandResponse message. Also converts values to other types if specified.
         * @param {containerlifecycle.StartCommandResponse} message StartCommandResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: containerlifecycle.StartCommandResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this StartCommandResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this StartCommandResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new StopContainerResponse.
     * @exports containerlifecycle.StopContainerResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class StopContainerResponse {

        /**
         * Constructs a new StopContainerResponse.
         * @exports containerlifecycle.StopContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * StopContainerResponse error.
         * @type {string|undefined}
         */
        error?: string;

        /**
         * Creates a new StopContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse instance
         */
        static create(properties?: Object): containerlifecycle.StopContainerResponse;

        /**
         * Encodes the specified StopContainerResponse message.
         * @param {containerlifecycle.StopContainerResponse|Object} message StopContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (containerlifecycle.StopContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified StopContainerResponse message, length delimited.
         * @param {containerlifecycle.StopContainerResponse|Object} message StopContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (containerlifecycle.StopContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StopContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): containerlifecycle.StopContainerResponse;

        /**
         * Decodes a StopContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): containerlifecycle.StopContainerResponse;

        /**
         * Verifies a StopContainerResponse message.
         * @param {containerlifecycle.StopContainerResponse|Object} message StopContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (containerlifecycle.StopContainerResponse|Object)): string;

        /**
         * Creates a StopContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        static fromObject(object: { [k: string]: any }): containerlifecycle.StopContainerResponse;

        /**
         * Creates a StopContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link containerlifecycle.StopContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {containerlifecycle.StopContainerResponse} StopContainerResponse
         */
        static from(object: { [k: string]: any }): containerlifecycle.StopContainerResponse;

        /**
         * Creates a plain object from a StopContainerResponse message. Also converts values to other types if specified.
         * @param {containerlifecycle.StopContainerResponse} message StopContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: containerlifecycle.StopContainerResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this StopContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this StopContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }
}

/**
 * Callback as used by {@link ContainerLifecycleService#startContainer}.
 * @typedef ContainerLifecycleService_startContainer_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {containerlifecycle.StartContainerResponse} [response] StartContainerResponse
 */
type ContainerLifecycleService_startContainer_Callback = (error: Error, response?: containerlifecycle.StartContainerResponse) => void;

/**
 * Callback as used by {@link ContainerLifecycleService#startCommand}.
 * @typedef ContainerLifecycleService_startCommand_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {containerlifecycle.StartCommandResponse} [response] StartCommandResponse
 */
type ContainerLifecycleService_startCommand_Callback = (error: Error, response?: containerlifecycle.StartCommandResponse) => void;

/**
 * Callback as used by {@link ContainerLifecycleService#stopContainer}.
 * @typedef ContainerLifecycleService_stopContainer_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {containerlifecycle.StopContainerResponse} [response] StopContainerResponse
 */
type ContainerLifecycleService_stopContainer_Callback = (error: Error, response?: containerlifecycle.StopContainerResponse) => void;

/**
 * Namespace customercontainer.
 * @exports customercontainer
 * @namespace
 */
export namespace customercontainer {

    /**
     * Constructs a new CustomerContainerService service.
     * @exports customercontainer.CustomerContainerService
     * @extends $protobuf.rpc.Service
     * @constructor
     * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
     * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
     * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
     */
    class CustomerContainerService extends $protobuf.rpc.Service {

        /**
         * Constructs a new CustomerContainerService service.
         * @exports customercontainer.CustomerContainerService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Creates new CustomerContainerService service using the specified rpc implementation.
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {CustomerContainerService} RPC service. Useful where requests and/or responses are streamed.
         */
        static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): CustomerContainerService;

        /**
         * Calls CreateContainer.
         * @param {customercontainer.CreateContainerRequest|Object} request CreateContainerRequest message or plain object
         * @param {CustomerContainerService_createContainer_Callback} callback Node-style callback called with the error, if any, and CreateContainerResponse
         * @returns {undefined}
         */
        createContainer(request: (customercontainer.CreateContainerRequest|Object), callback: CustomerContainerService_createContainer_Callback): void;

        /**
         * Calls EditContainer.
         * @param {customercontainer.EditContainerRequest|Object} request EditContainerRequest message or plain object
         * @param {CustomerContainerService_editContainer_Callback} callback Node-style callback called with the error, if any, and EditContainerResponse
         * @returns {undefined}
         */
        editContainer(request: (customercontainer.EditContainerRequest|Object), callback: CustomerContainerService_editContainer_Callback): void;

        /**
         * Calls RemoveContainer.
         * @param {customercontainer.RemoveContainerRequest|Object} request RemoveContainerRequest message or plain object
         * @param {CustomerContainerService_removeContainer_Callback} callback Node-style callback called with the error, if any, and RemoveContainerResponse
         * @returns {undefined}
         */
        removeContainer(request: (customercontainer.RemoveContainerRequest|Object), callback: CustomerContainerService_removeContainer_Callback): void;

        /**
         * Calls Instances.
         * @param {customercontainer.InstancesRequest|Object} request InstancesRequest message or plain object
         * @param {CustomerContainerService_instances_Callback} callback Node-style callback called with the error, if any, and InstancesResponse
         * @returns {undefined}
         */
        instances(request: (customercontainer.InstancesRequest|Object), callback: CustomerContainerService_instances_Callback): void;
    }

    /**
     * Constructs a new ContainerConfig.
     * @exports customercontainer.ContainerConfig
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class ContainerConfig {

        /**
         * Constructs a new ContainerConfig.
         * @exports customercontainer.ContainerConfig
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * ContainerConfig ImageName.
         * @type {string|undefined}
         */
        ImageName?: string;

        /**
         * Creates a new ContainerConfig instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.ContainerConfig} ContainerConfig instance
         */
        static create(properties?: Object): customercontainer.ContainerConfig;

        /**
         * Encodes the specified ContainerConfig message.
         * @param {customercontainer.ContainerConfig|Object} message ContainerConfig message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.ContainerConfig|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified ContainerConfig message, length delimited.
         * @param {customercontainer.ContainerConfig|Object} message ContainerConfig message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.ContainerConfig|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ContainerConfig message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.ContainerConfig;

        /**
         * Decodes a ContainerConfig message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.ContainerConfig;

        /**
         * Verifies a ContainerConfig message.
         * @param {customercontainer.ContainerConfig|Object} message ContainerConfig message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.ContainerConfig|Object)): string;

        /**
         * Creates a ContainerConfig message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        static fromObject(object: { [k: string]: any }): customercontainer.ContainerConfig;

        /**
         * Creates a ContainerConfig message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.ContainerConfig.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.ContainerConfig} ContainerConfig
         */
        static from(object: { [k: string]: any }): customercontainer.ContainerConfig;

        /**
         * Creates a plain object from a ContainerConfig message. Also converts values to other types if specified.
         * @param {customercontainer.ContainerConfig} message ContainerConfig
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.ContainerConfig, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this ContainerConfig message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this ContainerConfig to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new CreateContainerRequest.
     * @exports customercontainer.CreateContainerRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class CreateContainerRequest {

        /**
         * Constructs a new CreateContainerRequest.
         * @exports customercontainer.CreateContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * CreateContainerRequest Refid.
         * @type {number|undefined}
         */
        Refid?: number;

        /**
         * CreateContainerRequest Cfg.
         * @type {customercontainer.ContainerConfig|undefined}
         */
        Cfg?: customercontainer.ContainerConfig;

        /**
         * Creates a new CreateContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest instance
         */
        static create(properties?: Object): customercontainer.CreateContainerRequest;

        /**
         * Encodes the specified CreateContainerRequest message.
         * @param {customercontainer.CreateContainerRequest|Object} message CreateContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.CreateContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified CreateContainerRequest message, length delimited.
         * @param {customercontainer.CreateContainerRequest|Object} message CreateContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.CreateContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.CreateContainerRequest;

        /**
         * Decodes a CreateContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.CreateContainerRequest;

        /**
         * Verifies a CreateContainerRequest message.
         * @param {customercontainer.CreateContainerRequest|Object} message CreateContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.CreateContainerRequest|Object)): string;

        /**
         * Creates a CreateContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        static fromObject(object: { [k: string]: any }): customercontainer.CreateContainerRequest;

        /**
         * Creates a CreateContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerRequest} CreateContainerRequest
         */
        static from(object: { [k: string]: any }): customercontainer.CreateContainerRequest;

        /**
         * Creates a plain object from a CreateContainerRequest message. Also converts values to other types if specified.
         * @param {customercontainer.CreateContainerRequest} message CreateContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.CreateContainerRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this CreateContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this CreateContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new CreateContainerResponse.
     * @exports customercontainer.CreateContainerResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class CreateContainerResponse {

        /**
         * Constructs a new CreateContainerResponse.
         * @exports customercontainer.CreateContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * CreateContainerResponse Name.
         * @type {string|undefined}
         */
        Name?: string;

        /**
         * CreateContainerResponse ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * CreateContainerResponse Error.
         * @type {string|undefined}
         */
        Error?: string;

        /**
         * Creates a new CreateContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse instance
         */
        static create(properties?: Object): customercontainer.CreateContainerResponse;

        /**
         * Encodes the specified CreateContainerResponse message.
         * @param {customercontainer.CreateContainerResponse|Object} message CreateContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.CreateContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified CreateContainerResponse message, length delimited.
         * @param {customercontainer.CreateContainerResponse|Object} message CreateContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.CreateContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.CreateContainerResponse;

        /**
         * Decodes a CreateContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.CreateContainerResponse;

        /**
         * Verifies a CreateContainerResponse message.
         * @param {customercontainer.CreateContainerResponse|Object} message CreateContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.CreateContainerResponse|Object)): string;

        /**
         * Creates a CreateContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        static fromObject(object: { [k: string]: any }): customercontainer.CreateContainerResponse;

        /**
         * Creates a CreateContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateContainerResponse} CreateContainerResponse
         */
        static from(object: { [k: string]: any }): customercontainer.CreateContainerResponse;

        /**
         * Creates a plain object from a CreateContainerResponse message. Also converts values to other types if specified.
         * @param {customercontainer.CreateContainerResponse} message CreateContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.CreateContainerResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this CreateContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this CreateContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new EditContainerRequest.
     * @exports customercontainer.EditContainerRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class EditContainerRequest {

        /**
         * Constructs a new EditContainerRequest.
         * @exports customercontainer.EditContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * EditContainerRequest ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * EditContainerRequest Cfg.
         * @type {customercontainer.ContainerConfig|undefined}
         */
        Cfg?: customercontainer.ContainerConfig;

        /**
         * Creates a new EditContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest instance
         */
        static create(properties?: Object): customercontainer.EditContainerRequest;

        /**
         * Encodes the specified EditContainerRequest message.
         * @param {customercontainer.EditContainerRequest|Object} message EditContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.EditContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified EditContainerRequest message, length delimited.
         * @param {customercontainer.EditContainerRequest|Object} message EditContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.EditContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an EditContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.EditContainerRequest;

        /**
         * Decodes an EditContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.EditContainerRequest;

        /**
         * Verifies an EditContainerRequest message.
         * @param {customercontainer.EditContainerRequest|Object} message EditContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.EditContainerRequest|Object)): string;

        /**
         * Creates an EditContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        static fromObject(object: { [k: string]: any }): customercontainer.EditContainerRequest;

        /**
         * Creates an EditContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.EditContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerRequest} EditContainerRequest
         */
        static from(object: { [k: string]: any }): customercontainer.EditContainerRequest;

        /**
         * Creates a plain object from an EditContainerRequest message. Also converts values to other types if specified.
         * @param {customercontainer.EditContainerRequest} message EditContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.EditContainerRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this EditContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this EditContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new EditContainerResponse.
     * @exports customercontainer.EditContainerResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class EditContainerResponse {

        /**
         * Constructs a new EditContainerResponse.
         * @exports customercontainer.EditContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * EditContainerResponse Error.
         * @type {string|undefined}
         */
        Error?: string;

        /**
         * Creates a new EditContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse instance
         */
        static create(properties?: Object): customercontainer.EditContainerResponse;

        /**
         * Encodes the specified EditContainerResponse message.
         * @param {customercontainer.EditContainerResponse|Object} message EditContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.EditContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified EditContainerResponse message, length delimited.
         * @param {customercontainer.EditContainerResponse|Object} message EditContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.EditContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an EditContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.EditContainerResponse;

        /**
         * Decodes an EditContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.EditContainerResponse;

        /**
         * Verifies an EditContainerResponse message.
         * @param {customercontainer.EditContainerResponse|Object} message EditContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.EditContainerResponse|Object)): string;

        /**
         * Creates an EditContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        static fromObject(object: { [k: string]: any }): customercontainer.EditContainerResponse;

        /**
         * Creates an EditContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.EditContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.EditContainerResponse} EditContainerResponse
         */
        static from(object: { [k: string]: any }): customercontainer.EditContainerResponse;

        /**
         * Creates a plain object from an EditContainerResponse message. Also converts values to other types if specified.
         * @param {customercontainer.EditContainerResponse} message EditContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.EditContainerResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this EditContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this EditContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new RemoveContainerRequest.
     * @exports customercontainer.RemoveContainerRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class RemoveContainerRequest {

        /**
         * Constructs a new RemoveContainerRequest.
         * @exports customercontainer.RemoveContainerRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * RemoveContainerRequest ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * Creates a new RemoveContainerRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest instance
         */
        static create(properties?: Object): customercontainer.RemoveContainerRequest;

        /**
         * Encodes the specified RemoveContainerRequest message.
         * @param {customercontainer.RemoveContainerRequest|Object} message RemoveContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.RemoveContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified RemoveContainerRequest message, length delimited.
         * @param {customercontainer.RemoveContainerRequest|Object} message RemoveContainerRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.RemoveContainerRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a RemoveContainerRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.RemoveContainerRequest;

        /**
         * Decodes a RemoveContainerRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.RemoveContainerRequest;

        /**
         * Verifies a RemoveContainerRequest message.
         * @param {customercontainer.RemoveContainerRequest|Object} message RemoveContainerRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.RemoveContainerRequest|Object)): string;

        /**
         * Creates a RemoveContainerRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        static fromObject(object: { [k: string]: any }): customercontainer.RemoveContainerRequest;

        /**
         * Creates a RemoveContainerRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.RemoveContainerRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerRequest} RemoveContainerRequest
         */
        static from(object: { [k: string]: any }): customercontainer.RemoveContainerRequest;

        /**
         * Creates a plain object from a RemoveContainerRequest message. Also converts values to other types if specified.
         * @param {customercontainer.RemoveContainerRequest} message RemoveContainerRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.RemoveContainerRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this RemoveContainerRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this RemoveContainerRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new RemoveContainerResponse.
     * @exports customercontainer.RemoveContainerResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class RemoveContainerResponse {

        /**
         * Constructs a new RemoveContainerResponse.
         * @exports customercontainer.RemoveContainerResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * RemoveContainerResponse Error.
         * @type {string|undefined}
         */
        Error?: string;

        /**
         * Creates a new RemoveContainerResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse instance
         */
        static create(properties?: Object): customercontainer.RemoveContainerResponse;

        /**
         * Encodes the specified RemoveContainerResponse message.
         * @param {customercontainer.RemoveContainerResponse|Object} message RemoveContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.RemoveContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified RemoveContainerResponse message, length delimited.
         * @param {customercontainer.RemoveContainerResponse|Object} message RemoveContainerResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.RemoveContainerResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a RemoveContainerResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.RemoveContainerResponse;

        /**
         * Decodes a RemoveContainerResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.RemoveContainerResponse;

        /**
         * Verifies a RemoveContainerResponse message.
         * @param {customercontainer.RemoveContainerResponse|Object} message RemoveContainerResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.RemoveContainerResponse|Object)): string;

        /**
         * Creates a RemoveContainerResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        static fromObject(object: { [k: string]: any }): customercontainer.RemoveContainerResponse;

        /**
         * Creates a RemoveContainerResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.RemoveContainerResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.RemoveContainerResponse} RemoveContainerResponse
         */
        static from(object: { [k: string]: any }): customercontainer.RemoveContainerResponse;

        /**
         * Creates a plain object from a RemoveContainerResponse message. Also converts values to other types if specified.
         * @param {customercontainer.RemoveContainerResponse} message RemoveContainerResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.RemoveContainerResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this RemoveContainerResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this RemoveContainerResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new InstancesRequest.
     * @exports customercontainer.InstancesRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class InstancesRequest {

        /**
         * Constructs a new InstancesRequest.
         * @exports customercontainer.InstancesRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * InstancesRequest Refid.
         * @type {number|undefined}
         */
        Refid?: number;

        /**
         * Creates a new InstancesRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.InstancesRequest} InstancesRequest instance
         */
        static create(properties?: Object): customercontainer.InstancesRequest;

        /**
         * Encodes the specified InstancesRequest message.
         * @param {customercontainer.InstancesRequest|Object} message InstancesRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.InstancesRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified InstancesRequest message, length delimited.
         * @param {customercontainer.InstancesRequest|Object} message InstancesRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.InstancesRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InstancesRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.InstancesRequest;

        /**
         * Decodes an InstancesRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.InstancesRequest;

        /**
         * Verifies an InstancesRequest message.
         * @param {customercontainer.InstancesRequest|Object} message InstancesRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.InstancesRequest|Object)): string;

        /**
         * Creates an InstancesRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        static fromObject(object: { [k: string]: any }): customercontainer.InstancesRequest;

        /**
         * Creates an InstancesRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.InstancesRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesRequest} InstancesRequest
         */
        static from(object: { [k: string]: any }): customercontainer.InstancesRequest;

        /**
         * Creates a plain object from an InstancesRequest message. Also converts values to other types if specified.
         * @param {customercontainer.InstancesRequest} message InstancesRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.InstancesRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this InstancesRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this InstancesRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new InstancesResponse.
     * @exports customercontainer.InstancesResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class InstancesResponse {

        /**
         * Constructs a new InstancesResponse.
         * @exports customercontainer.InstancesResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * InstancesResponse Instances.
         * @type {Array.<string>|undefined}
         */
        Instances?: string[];

        /**
         * Creates a new InstancesResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.InstancesResponse} InstancesResponse instance
         */
        static create(properties?: Object): customercontainer.InstancesResponse;

        /**
         * Encodes the specified InstancesResponse message.
         * @param {customercontainer.InstancesResponse|Object} message InstancesResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.InstancesResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified InstancesResponse message, length delimited.
         * @param {customercontainer.InstancesResponse|Object} message InstancesResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.InstancesResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InstancesResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.InstancesResponse;

        /**
         * Decodes an InstancesResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.InstancesResponse;

        /**
         * Verifies an InstancesResponse message.
         * @param {customercontainer.InstancesResponse|Object} message InstancesResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.InstancesResponse|Object)): string;

        /**
         * Creates an InstancesResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        static fromObject(object: { [k: string]: any }): customercontainer.InstancesResponse;

        /**
         * Creates an InstancesResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.InstancesResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.InstancesResponse} InstancesResponse
         */
        static from(object: { [k: string]: any }): customercontainer.InstancesResponse;

        /**
         * Creates a plain object from an InstancesResponse message. Also converts values to other types if specified.
         * @param {customercontainer.InstancesResponse} message InstancesResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.InstancesResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this InstancesResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this InstancesResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new CreateDockerImageRequest.
     * @exports customercontainer.CreateDockerImageRequest
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class CreateDockerImageRequest {

        /**
         * Constructs a new CreateDockerImageRequest.
         * @exports customercontainer.CreateDockerImageRequest
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * CreateDockerImageRequest Refid.
         * @type {number|undefined}
         */
        Refid?: number;

        /**
         * CreateDockerImageRequest KmiID.
         * @type {number|undefined}
         */
        KmiID?: number;

        /**
         * Creates a new CreateDockerImageRequest instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest instance
         */
        static create(properties?: Object): customercontainer.CreateDockerImageRequest;

        /**
         * Encodes the specified CreateDockerImageRequest message.
         * @param {customercontainer.CreateDockerImageRequest|Object} message CreateDockerImageRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.CreateDockerImageRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified CreateDockerImageRequest message, length delimited.
         * @param {customercontainer.CreateDockerImageRequest|Object} message CreateDockerImageRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.CreateDockerImageRequest|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateDockerImageRequest message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.CreateDockerImageRequest;

        /**
         * Decodes a CreateDockerImageRequest message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.CreateDockerImageRequest;

        /**
         * Verifies a CreateDockerImageRequest message.
         * @param {customercontainer.CreateDockerImageRequest|Object} message CreateDockerImageRequest message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.CreateDockerImageRequest|Object)): string;

        /**
         * Creates a CreateDockerImageRequest message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        static fromObject(object: { [k: string]: any }): customercontainer.CreateDockerImageRequest;

        /**
         * Creates a CreateDockerImageRequest message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateDockerImageRequest.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageRequest} CreateDockerImageRequest
         */
        static from(object: { [k: string]: any }): customercontainer.CreateDockerImageRequest;

        /**
         * Creates a plain object from a CreateDockerImageRequest message. Also converts values to other types if specified.
         * @param {customercontainer.CreateDockerImageRequest} message CreateDockerImageRequest
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.CreateDockerImageRequest, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this CreateDockerImageRequest message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this CreateDockerImageRequest to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }

    /**
     * Constructs a new CreateDockerImageResponse.
     * @exports customercontainer.CreateDockerImageResponse
     * @constructor
     * @param {Object} [properties] Properties to set
     */
    class CreateDockerImageResponse {

        /**
         * Constructs a new CreateDockerImageResponse.
         * @exports customercontainer.CreateDockerImageResponse
         * @constructor
         * @param {Object} [properties] Properties to set
         */
        constructor(properties?: Object);

        /**
         * CreateDockerImageResponse ID.
         * @type {string|undefined}
         */
        ID?: string;

        /**
         * CreateDockerImageResponse Error.
         * @type {string|undefined}
         */
        Error?: string;

        /**
         * Creates a new CreateDockerImageResponse instance using the specified properties.
         * @param {Object} [properties] Properties to set
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse instance
         */
        static create(properties?: Object): customercontainer.CreateDockerImageResponse;

        /**
         * Encodes the specified CreateDockerImageResponse message.
         * @param {customercontainer.CreateDockerImageResponse|Object} message CreateDockerImageResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encode(message: (customercontainer.CreateDockerImageResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified CreateDockerImageResponse message, length delimited.
         * @param {customercontainer.CreateDockerImageResponse|Object} message CreateDockerImageResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        static encodeDelimited(message: (customercontainer.CreateDockerImageResponse|Object), writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateDockerImageResponse message from the specified reader or buffer.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): customercontainer.CreateDockerImageResponse;

        /**
         * Decodes a CreateDockerImageResponse message from the specified reader or buffer, length delimited.
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): customercontainer.CreateDockerImageResponse;

        /**
         * Verifies a CreateDockerImageResponse message.
         * @param {customercontainer.CreateDockerImageResponse|Object} message CreateDockerImageResponse message or plain object to verify
         * @returns {?string} `null` if valid, otherwise the reason why it is not
         */
        static verify(message: (customercontainer.CreateDockerImageResponse|Object)): string;

        /**
         * Creates a CreateDockerImageResponse message from a plain object. Also converts values to their respective internal types.
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        static fromObject(object: { [k: string]: any }): customercontainer.CreateDockerImageResponse;

        /**
         * Creates a CreateDockerImageResponse message from a plain object. Also converts values to their respective internal types.
         * This is an alias of {@link customercontainer.CreateDockerImageResponse.fromObject}.
         * @function
         * @param {Object.<string,*>} object Plain object
         * @returns {customercontainer.CreateDockerImageResponse} CreateDockerImageResponse
         */
        static from(object: { [k: string]: any }): customercontainer.CreateDockerImageResponse;

        /**
         * Creates a plain object from a CreateDockerImageResponse message. Also converts values to other types if specified.
         * @param {customercontainer.CreateDockerImageResponse} message CreateDockerImageResponse
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        static toObject(message: customercontainer.CreateDockerImageResponse, options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Creates a plain object from this CreateDockerImageResponse message. Also converts values to other types if specified.
         * @param {$protobuf.ConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        toObject(options?: $protobuf.ConversionOptions): { [k: string]: any };

        /**
         * Converts this CreateDockerImageResponse to JSON.
         * @returns {Object.<string,*>} JSON object
         */
        toJSON(): { [k: string]: any };
    }
}

/**
 * Callback as used by {@link CustomerContainerService#createContainer}.
 * @typedef CustomerContainerService_createContainer_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {customercontainer.CreateContainerResponse} [response] CreateContainerResponse
 */
type CustomerContainerService_createContainer_Callback = (error: Error, response?: customercontainer.CreateContainerResponse) => void;

/**
 * Callback as used by {@link CustomerContainerService#editContainer}.
 * @typedef CustomerContainerService_editContainer_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {customercontainer.EditContainerResponse} [response] EditContainerResponse
 */
type CustomerContainerService_editContainer_Callback = (error: Error, response?: customercontainer.EditContainerResponse) => void;

/**
 * Callback as used by {@link CustomerContainerService#removeContainer}.
 * @typedef CustomerContainerService_removeContainer_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {customercontainer.RemoveContainerResponse} [response] RemoveContainerResponse
 */
type CustomerContainerService_removeContainer_Callback = (error: Error, response?: customercontainer.RemoveContainerResponse) => void;

/**
 * Callback as used by {@link CustomerContainerService#instances}.
 * @typedef CustomerContainerService_instances_Callback
 * @type {function}
 * @param {?Error} error Error, if any
 * @param {customercontainer.InstancesResponse} [response] InstancesResponse
 */
type CustomerContainerService_instances_Callback = (error: Error, response?: customercontainer.InstancesResponse) => void;
