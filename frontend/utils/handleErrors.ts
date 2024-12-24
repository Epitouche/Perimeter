import type { ErrorResponse } from "~/interfaces/error";

export function handleError(error: unknown): never {
  let statusCode = 500;
  let message = "An unknown error occurred";

  if (error instanceof Error) {
    if ("statusCode" in error) {
      statusCode = (error as ErrorResponse).statusCode || statusCode;
    }
    message = error.message;
  } else if (typeof error === "object" && error !== null && "statusCode" in error) {
    statusCode = (error as ErrorResponse).statusCode || statusCode;
  }

  throw createError({
    statusCode,
    message,
  });
}
