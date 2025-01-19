import type { ErrorResponse } from "~/interfaces/error";

/**
 * This function creates an error object that can be thrown
 * 
 * @param error - The error to handle
 */
export function handleError(error: unknown): never {
  let statusCode = 500;
  let message = "An unknown error occurred";

  if (error instanceof Error) {
    if ("statusCode" in error) {
      statusCode = (error as ErrorResponse).statusCode || statusCode;
    }
    message = error.message;
  } else if (
    typeof error === "object" &&
    error !== null &&
    "statusCode" in error
  ) {
    statusCode = (error as ErrorResponse).statusCode || statusCode;
  }

  throw createError({
    statusCode,
    message,
  });
}
