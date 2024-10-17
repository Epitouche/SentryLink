/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{vue,js,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
  ],
  theme: {
    extend: {
      colors: {
        cprimary: {
          light: "#f0f4f8", // Gris clair
          DEFAULT: "#1a202c", // Gris foncé
          dark: "#000000", // Noir
        },
        secondary: {
          light: "#d1fae5", // Vert clair
          DEFAULT: "#059669", // Vert
          dark: "#064e3b", // Vert foncé
        },
        accent: {
          DEFAULT: "#f97316", // Orange
        },
      },
    },
  },
  plugins: [],
};
