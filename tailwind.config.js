module.exports = {
  content: [
    "web/view/**/*.templ",
    "web/static/tailwind/node_modules/tw-elements/dist/js/**/*.js"
  ],
  plugins: [require("./web/static/tailwind/node_modules/tw-elements/dist/plugin.cjs")],
  darkMode: "class"
};
//@npx tailwindcss build -i ../styles/styles.css -o ../styles/output.css
// @npx tailwindcss build -i web/static/styles/styles.css -o web/static/styles/output.css --minify