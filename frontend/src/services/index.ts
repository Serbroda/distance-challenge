import type {ModelsUser} from "../swagger";
import {AuthStore} from "./auth/authstore";

const {VITE_BACKEND_BASE_URL} = import.meta.env;

const basePath: string = VITE_BACKEND_BASE_URL || "/";

const authStore = new AuthStore<ModelsUser>();
