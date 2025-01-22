/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue",
    "./node_modules/flowbite/**/*.{js,ts}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "miku-bg1": "url('')",
        luotianyi: "url('/images/background/luotianyi.jpg')",
      },
    },
  },
  plugins: [require("flowbite/plugin")],
};
