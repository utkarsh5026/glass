import { configureStore } from "@reduxjs/toolkit";
import mentorReducer from "./people/mentorSlice";
import studentReducer from "./people/studentSlice";
import authReducer from "./auth/authSlice";
import assignmentReducer from "./assignments/slice";
import dashboardReducer from "./dasboard/slice";
import coursesReducer from "./courses/slice";
import materialReducer from "./materials/slice";

const store = configureStore({
  reducer: {
    mentors: mentorReducer,
    students: studentReducer,
    assignments: assignmentReducer,
    auth: authReducer,
    dashboard: dashboardReducer,
    courses: coursesReducer,
    materials: materialReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
