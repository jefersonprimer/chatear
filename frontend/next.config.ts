import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  transpilePackages: ["apollo-upload-client"],
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "res.cloudinary.com",
        port: "",
        pathname: "/**",
      },
    ],
  },
};

export default nextConfig;
