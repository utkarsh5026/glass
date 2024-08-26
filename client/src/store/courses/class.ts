import {createSlice} from "@reduxjs/toolkit";
import {CourseBasic} from "./types";


interface ClassState {
    classes: CourseBasic[];
}

const initialState: ClassState = {
    classes: []
};

const classSlice = createSlice({
    name: 'classes',
    initialState,
    reducers: {
        addClass(state, action) {
            state.classes.push(action.payload);
        }
    }
});


export const {addClass} = classSlice.actions;
export default classSlice.reducer;