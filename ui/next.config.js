/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "export",
    // reactStrictMode: true,
    env: {
        // apiUrl: "http://localhost:8080/api",
        // apiUrl: "https://mypetjournal-lqkz3.ondigitalocean.app/api",
        apiUrl: "/api",
    },
    images: {
        unoptimized: true,
    },
    distDir: "build",
}

module.exports = nextConfig
