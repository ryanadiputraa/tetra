import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        primary: "#4C758A",
        "primary-dark": "#759EB3",
        secondary: "#54A6CF",
        accent: "#BADFF2",
      },
    },
  },
  plugins: [],
} satisfies Config;
