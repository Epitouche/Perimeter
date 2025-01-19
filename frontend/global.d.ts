// global.d.ts
import { defineNuxtRouteMiddleware, navigateTo, useCookie } from 'nuxt/app';

declare global {
  const defineNuxtRouteMiddleware: (...args: any[]) => any;
  const navigateTo: (...args: any[]) => any;
  const useCookie: (...args: any[]) => any;
  const defineEventHandler: (...args: any[]) => any;
  const readBody: (...args: any[]) => any;
  const createError: (...args: any[]) => any;
  const defineStore: (...args: any[]) => any;
  const handleErrorStatus: (...args: any[]) => any;
}