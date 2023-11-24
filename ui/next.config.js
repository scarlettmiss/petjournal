/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "export",
    // reactStrictMode: true,
    env: {
        apiUrl: "https://petjournal-app-aeqg2.ondigitalocean.app/api",
        // apiUrl: "http://localhost:8080/api",
    },
    images: {
        unoptimized: true
    },
    distDir: 'build'
}

module.exports = nextConfig
