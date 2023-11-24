module.exports = {
    root: true,
    extends: ["next/core-web-vitals"],
    env: {
        jest: true,
    },
    rules: {
        "@typescript-eslint/no-unused-vars": "off",
        "import/extensions": "off",
        "import/no-extraneous-dependencies": "off",
        "@typescript-eslint/lines-between-class-members": "off",
        "@typescript-eslint/no-use-before-define": "off",
        "@typescript-eslint/no-array-constructor": "off",
        "@typescript-eslint/no-shadow": "off",
        "no-script-url": "off",
        "@typescript-eslint/default-param-last": "off",
        "react/no-did-mount-set-state": "off",
        "react/no-did-update-set-state": "off",
        radix: "off",
        "no-bitwise": "off",
        "no-control-regex": "off",
        "script-eslint/no-unused-expressions": "off",
        "@typescript-eslint/return-await": "off",
        "@typescript-eslint/dot-notation": "off",
        "@typescript-eslint/no-loop-func": "off",
        "@typescript-eslint/no-unused-expressions": "off",
        "no-useless-escape": "off",
        "prefer-const": ["error", {destructuring: "all", ignoreReadBeforeAssign: true}],
        "default-case": "off",
        eqeqeq: "off",
        "react-native/no-inline-styles": "off",
    },
    parserOptions: {
        project: "tsconfig.json",
        tsconfigRootDir: __dirname,
    },
    ignorePatterns: ["html/**", "android/**", "ios/**", ".eslintrc.js"],
}
