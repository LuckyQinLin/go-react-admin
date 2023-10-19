import {create} from "zustand";
import {User} from "@/types";
import {createJSONStorage, devtools, persist} from "zustand/middleware";
import {userLogin, userLoginInfo} from "@/api/user.ts";

const useStore = create<User.UserStoreProp>()(
    devtools(
        persist(
            set => ({
                setLoginProp: (loginProp) => set({ loginProp }),
                setUserProp: (userProp) => set({ userProp }),
                userLoginFetch: async (param: User.LoginFormProp) => {
                    const response = await userLogin(param);
                    console.log("login", response);
                    set({ loginProp: response })
                },
                useInfoFetch: async () => {
                    const response = await userLoginInfo()
                    set({userProp: response})
                }
            }),
            {
                name: 'admin-storage',
                storage: createJSONStorage(() => sessionStorage)
            }
        )
    )
)


export default useStore;
