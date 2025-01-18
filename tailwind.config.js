/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.html"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["system-ui", "sans-serif"],
        mono: ["system-ui", "monospace"],
        serif: ["system-ui", "serif"],
      },
    },
  },
  plugins: [],
  safelist: [
    'hidden',
    'invisible',
    'opactity-0',
  ]
};

