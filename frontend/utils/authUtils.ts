// ~/utils/authUtils.ts
import type { Ref } from "vue";
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import type { OAuthLink } from "~/interfaces/authLink";
import type { Token } from "~/interfaces/serviceResponse";

/**
 * Fetches the data from the API with the link provided.
 * 
 * @param label - The label of the service.
 * @returns - Returns the response.
 */
export const authApiCall = async (label: string) => {
  try {
    const response = await $fetch<OAuthLink>("/api/auth/service/redirect", {
      method: "POST",
      body: {
        link: label,
      },
    });
    navigateTo(response.authentication_url, { external: true });
    return response;
  } catch (error: unknown) {
    throw handleErrorStatus(error);
  }
};

/**
 * Disconnects the service from the user.
 * 
 * @param token - The token of the user.
 * @param tokenId - The id of the token.
 * @returns - Returns the response.
 */
export const disconnectService = async (token: string, tokenId: number) => {
  try {
    const response = await $fetch("/api/auth/service/disconnection", {
      method: "POST",
      body: {
        authorization: token,
        tokenId: tokenId,
      },
    });
    return response;
  } catch (error: unknown) {
    throw handleErrorStatus(error);
  }
};

/**
 * Handles the click event of the service.
 * 
 * @param label - The label of the service.
 * @param services - The services of the user.
 * @param tokens - The tokens of the user.
 * @param token - The token of the user.
 * @returns - Returns true if the service is disconnected, false otherwise.
 */
export const handleClick = async (
  label: string,
  services: Ref<ServiceInfo[]>,
  tokens?: Ref<Token[]>,
  token?: string
) => {
  const serviceNames = services.value.map((service) => service.name);
  let matchingToken;
  if (tokens && token) {
    matchingToken = tokens.value.find((t) => t.service.name === label);
  }
  if (matchingToken && token) {
    const serviceId = matchingToken.id;
    await disconnectService(token, Number(serviceId));
    return true;
  } else {
    const apiLink = `http://server:8080/api/v1/${label.toLowerCase()}/auth/`;

    if (
      serviceNames.includes(label) &&
      label.toLowerCase() != "timer" &&
      label.toLowerCase() != "openweathermap"
    ) {
      await authApiCall(apiLink);
      return false;
    } else {
      console.error(`Unknown service "${label}" clicked.`);
      return false;
    }
  }
};
