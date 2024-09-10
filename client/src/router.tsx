import { createBrowserRouter } from "react-router-dom";
import { Suspense } from "react";
import { Spin } from "antd";
import Layout from "./Layout";
import Auth from "./components/auth/Auth";
import NotFound from "./components/error/NotFound";
import Dashboard from "./components/dashboard/Dashboard";
import UserCourses from "./components/courses/list/UserCourses";
import CourseOverview from "./components/courses/CourseOverview";
import CreateCourseComponent from "./components/courses/create/CreateCourseComponent";

const SuspenseWrapper = ({ children }: { children: React.ReactNode }) => (
  <Suspense fallback={<Spin size="large" />}>{children}</Suspense>
);

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        path: "",
        element: (
          <SuspenseWrapper>
            <Dashboard />
          </SuspenseWrapper>
        ),
      },
      {
        path: "/courses",
        element: (
          <SuspenseWrapper>
            <UserCourses />
          </SuspenseWrapper>
        ),
      },
      {
        path: "/courses/create",
        element: (
          <SuspenseWrapper>
            <CreateCourseComponent />
          </SuspenseWrapper>
        ),
      },
      {
        path: "/courses/:courseId",
        element: (
          <SuspenseWrapper>
            <CourseOverview />
          </SuspenseWrapper>
        ),
      },
    ],
  },
  {
    path: "/login",
    element: (
      <SuspenseWrapper>
        <Auth />
      </SuspenseWrapper>
    ),
  },
  {
    path: "*",
    element: <NotFound />,
  },
]);
