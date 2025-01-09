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
    console.log("Response of authApiCall:", response);
    navigateTo(response.authentication_url, { external: true });
    return response;
  } catch (error: unknown) {
    throw handleErrorStatus(error);
  }
};

export const disconnectService = async (tokenId: string, token: string) => {
  try {
    const response = await $fetch("/api/auth/service/disconnection", {
      method: "POST",
      body: {
        token: token,
        tokenId: tokenId,
      },
    });
    console.log("Response of disconnectService:", response);
    return response;
  } catch (error: unknown) {
    throw handleErrorStatus(error);
  }
};

export const handleClick = (
  label: string,
  services: Ref<ServiceInfo[]>,
  tokens: Ref<Token[]>,
  token: string,
) => {
  const serviceNames = services.value.map((service) => service.name);

  // Trouver le tokenId correspondant au service
  const matchingToken = tokens.value.find((t) => t.service.name === label);

  if (matchingToken) {
    console.log(`Disconnecting service: ${label}`);
    disconnectService(matchingToken, token); // Envoie tokenId et token
  } else {
    const apiLink = `http://server:8080/api/v1/${label.toLowerCase()}/auth/`;

    if (serviceNames.includes(label)) {
      console.log(`Service "${label}" exists, initiating auth API call.`);
      authApiCall(apiLink);
    } else {
      console.log(`Unknown service "${label}" clicked.`);
    }
  }
};
