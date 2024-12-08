export default defineAppConfig({
  ui: {
    strategy: "override",
    button: {
            color: {
                black: {
                    solid: 'shadow-sm text-white dark:text-gray-900 bg-gray-900 dark:bg-white focus-visible:ring-inset focus-visible:ring-2 focus-visible:ring-primary-500 dark:focus-visible:ring-primary-400 disabled:opacity-100 disabled:bg-white disabled:text-black aria-disabled:opacity-100 aria-disabled:bg-white aria-disabled:text-black hover:none'
                }
            },
      default: {
        color: "black",
      },
      rounded: "rounded-custom_border_radius",
    },
    input: {
      base: "w-full focus:outline !border-custom_border_width border-custom_color-border opacity-100",
      rounded: "rounded-custom_border_radius",
      placeholder: "!px-5",
      color: {
        white: {
          outline: "shadow-none bg-custom_color-input ring-0",
        },
      },
    },
    container: {
      base: "mx-0",
      constrained: "max-w-[90%]",
    },
  },
});
