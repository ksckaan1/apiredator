/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        "default-bg": "#0D0F18",
        "accent-bg": "#1D2027",
        "primary": "#E1AF4B",
        "on-primary": "#734D00"
      }
    },
  },
  plugins: [],
}

