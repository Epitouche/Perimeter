// ~/utils/servicesConnectionInfos.ts
import { handleErrorStatus } from "./handleErrorStatus";
import type { ServiceResponse } from "~/interfaces/serviceResponse";

/**
 * Fetches the service connection information.
 * 
 * @param token - User token
 * @returns - Service connection information
 */
export async function servicesConnectionInfos(
  token: string
): Promise<ServiceResponse> {
  try {
    const response = await $fetch<ServiceResponse>("/api/auth/service/infos", {
      method: "POST",
      body: {
        authorization: token,
      },
    });

    if (response) {
      return response;
    } else {
      console.error("Response does not contain valid service connection data.");
      throw new Error("Invalid service connection data.");
    }
  } catch (error: unknown) {
    const errorMessage = handleErrorStatus(error);
    throw new Error(errorMessage || "Failed to fetch service connection info.");
  }
}
