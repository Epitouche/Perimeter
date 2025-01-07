// ~/utils/authUtils.ts
import type { Ref } from 'vue';
import type { ServiceInfo } from '~/interfaces/serviceinfo';
import type { OAuthLink } from '~/interfaces/authLink';

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

export const handleClick = (
  label: string,
  services: Ref<ServiceInfo[]>,
  serviceConnected?: Ref<string[]>,
) => {
  const serviceNames = services.value.map((service) => service.name);
  
  if (serviceConnected && serviceConnected.value.includes(label)) {
    //disconnectService(label);
  } else {
    const apiLink = `http://server:8080/api/v1/${label.toLowerCase()}/auth/`;
    console.log("apiLink:", apiLink);

    if (serviceNames.includes(label)) {
      console.log("serviceNames have label, sending to authApiCall");
      authApiCall(apiLink);
    } else {
      console.log(`${label} unknown icon clicked`);
    }
  }
};
