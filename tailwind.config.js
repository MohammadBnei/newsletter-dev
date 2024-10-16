/** @type {import('tailwindcss').Config} */
module.exports = {
  relative: true,
  content: ["./domain/template/**.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
