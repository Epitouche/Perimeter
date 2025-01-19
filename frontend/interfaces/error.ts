/**
 * Represents the error response with the following properties:
 * @property {number} statusCode: The status code of the error response.
 * @property {string} message: The error message.
 */
export interface ErrorResponse {
  statusCode?: number;
  message: string;
}
