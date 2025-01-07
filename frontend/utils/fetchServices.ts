import type { ServiceInfo } from "~/interfaces/serviceinfo";
import { handleErrorStatus } from "./handleErrorStatus";

export const fetchServices = async (): Promise<ServiceInfo[]> => {
  try {
    const result = await $fetch<ServiceInfo[]>("/api/workflow/services", {
      method: "GET",
    });
    return result;
  } catch (error) {
    throw handleErrorStatus(error);
  }
};