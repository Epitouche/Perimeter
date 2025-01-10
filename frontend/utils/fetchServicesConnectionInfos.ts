// ~/utils/servicesConnectionInfos.ts
import { handleErrorStatus } from "./handleErrorStatus";
import type { Token, ServiceResponse } from "~/interfaces/serviceResponse";

export async function servicesConnectionInfos(token: string): Promise<Token[]> {
  try {
    const response = await $fetch<ServiceResponse>("/api/auth/service/infos", {
      method: "POST",
      body: {
        authorization: token,
      },
    });

    if (response && Array.isArray(response.tokens)) {
      return response.tokens;
    } else {
      console.error("Response does not contain valid tokens.");
      return [];
    }
  } catch (error: unknown) {
    const errorMessage = handleErrorStatus(error);
    if (errorMessage === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    throw new Error(errorMessage || "Failed to fetch service connection info.");
  }
}
