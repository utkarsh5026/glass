import { createBrowserRouter } from "react-router-dom";
import Auth from "./components/auth/Auth";
import NotFound from "./components/error/NotFound";
import Dashboard from "./components/dashboard/Dashboard";
import UserCourses from "./components/courses/UserCourses";
import { Suspense } from "react";
import { Spin } from "antd";

export const router = createBrowserRouter([
  {
    path: "/login",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <Auth />
      </Suspense>
    ),
  },
  {
    path: "",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <Dashboard />
      </Suspense>
    ),
  },
  {
    path: "/courses",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <UserCourses />
      </Suspense>
    ),
  },
  {
    path: "*",
    element: <NotFound />,
  },
]);
