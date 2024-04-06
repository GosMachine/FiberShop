module.exports = {
  content: [
    "../../templates/**/*.{html,js}",
    "./node_modules/tw-elements/dist/js/**/*.js"
  ],
  plugins: [require("tw-elements/dist/plugin.cjs")],
  darkMode: "class"
};
//@npx tailwindcss build ../styles/styles.css -o ../styles/output.css