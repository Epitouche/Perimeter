import { describe, it, expect } from 'vitest';
import { mount } from '@vue/test-utils';
import { setup } from '@nuxt/test-utils';

import BackButton from '../components/BackButton.vue';

describe('BackButton', () => {
  setup();

  it('renders with the correct classes when isWhite is false', () => {
    const wrapper = mount(BackButton, {
      props: {
        link: '/dashboard',
        isWhite: false,
      },
    });

    expect(wrapper.find('UButton').classes()).toContain('!border-black');
    expect(wrapper.find('UButton').classes()).toContain('text-black');
  });

  it('renders with the correct classes when isWhite is true', () => {
    const wrapper = mount(BackButton, {
      props: {
        link: '/dashboard',
        isWhite: true,
      },
    });

    expect(wrapper.find('UButton').classes()).toContain('!border-white');
    expect(wrapper.find('UButton').classes()).toContain('text-white');
  });

  it('renders the link correctly', () => {
    const wrapper = mount(BackButton, {
      props: {
        link: '/dashboard',
        isWhite: false,
      },
    });

    expect(wrapper.find('UButton').attributes('to')).toBe('/dashboard');
  });
});
