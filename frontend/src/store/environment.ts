import { writable } from "svelte/store";

export type EnvVar = {
  key: string,
  value: string,
}

export const activeEnvVars = writable<EnvVar[]>([
  {
    key: "host",
    value: "https://kaanksc.com"
  },
  {
    key: "user_id",
    value: "123"
  },
  {
    key: "token",
    value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
]);