/**
 * Represents the response from the server when a user logs in.
 */

/**
 * @property {string} username - The username of the user.
 * @property {string} email - The email of the user.
 * @property {number} id - The id of the user.
 */
export interface User {
  username: string;
  email: string;
  id: number;
}

/**
 * @property {string} name - The name of the service.
 * @property {boolean} oauth - Whether the service uses OAuth.
 */
export interface Service {
  name: string;
  oauth: boolean;
}

/**
 * @property {Service} service - The service the token is for.
 * @property {number} id - The id of the token.
 */
export interface Token {
  service: Service;
  id: number;
}

/**
 * @property {Token[]} tokens - The tokens for the user.
 * @property {User} user - The user.
 */
export interface ServiceResponse {
  tokens: Token[];
  user: User;
}
