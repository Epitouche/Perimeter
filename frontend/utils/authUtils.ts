// ~/utils/authUtils.ts
import type { Ref } from "vue";
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import type { OAuthLink } from "~/interfaces/authLink";
import type { Token } from "~/interfaces/serviceResponse";

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
