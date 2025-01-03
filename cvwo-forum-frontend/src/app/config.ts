type AppConfig = {
  backendUrl: string;
  nodeEnv: "development" | "production" | "test";
};
export const APP_CONFIG: AppConfig = {
  backendUrl: process.env.BACKEND_URL || "http://127.0.0.1:5000",
  nodeEnv: process.env.NODE_ENV,
};
