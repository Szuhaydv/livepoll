/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontFamily: {
        actionJackson: ['Action Jackson', 'sans-serif'],
        caveatBrush: ['Caveat Brush', 'sans-serif'],
        bangers: ['Bangers', 'sans-serif'],
      }
    },
  },
  plugins: [],
}

