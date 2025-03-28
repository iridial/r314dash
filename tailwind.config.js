/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.html"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["system-ui", "sans-serif"],
        mono: ["ui-monospace", "monospace"],
        serif: ["ui-serif", "serif"],
      },
    },
  },
  //darkMode: 'media',
  darkMode: 'class',
  plugins: [],
  safelist: [
    'hidden',
    'invisible',
    'opactity-0',
  ]
};

