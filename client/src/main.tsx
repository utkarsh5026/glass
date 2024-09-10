import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ConfigProvider, Layout } from "antd";
import "./index.css";
import { Provider } from "react-redux";
import store from "./store/store.ts";
import { router } from "./router.tsx";
import { RouterProvider } from "react-router-dom";
import AppBar from "./AppBar";

const { Header, Content } = Layout;

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ConfigProvider
      theme={{
        token: {
          fontFamily: "'Cascadia Code', monospace",
          colorPrimary: "#1890ff",
          colorSuccess: "#52c41a",
          colorWarning: "#faad14",
          colorError: "#f5222d",
          colorTextBase: "#000000",
          colorBgBase: "#ffffff",
          borderRadius: 8,
        },
        components: {
          Typography: {
            colorPrimary: "#ff7",
          },
        },
      }}
    >
      <Provider store={store}>
        <Layout>
          <Header style={{ padding: 0, background: "#fff" }}>
            <AppBar />
          </Header>
          <Content>
            <RouterProvider router={router} />
          </Content>
        </Layout>
      </Provider>
    </ConfigProvider>
  </StrictMode>
);
