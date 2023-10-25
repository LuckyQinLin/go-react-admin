import {create} from "zustand";
import {User} from "@/types";
import {createJSONStorage, devtools, persist} from "zustand/middleware";
import {userLogin, userLoginInfo} from "@/api/user.ts";

const useStore = create<User.UserStoreProp>()(
    devtools(
        persist(
            (set, get) => ({
                tabViewKey: '/home/index',
                tabViews: [{key: '/home/index', title: '首页', closeIcon: true}],
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
                },
                addTabView: (data) => {
                    if (get().tabViews.findIndex(item => item.key == data.key) === -1) {
                        set((state) => ({tabViews: [...state.tabViews, data]}))
                    }
                    set({tabViewKey: data.key as string})
                },
                // 移除tabView
                removeTabView: (key, isNegation) => {
                    let target =  [...get().tabViews];
                    if (isNegation) { // 删除除了自身以及首页
                        const newData = target.filter(item => item.key === '/home/index' || item.key === key);
                        if (get().tabViewKey === key) {
                            set({tabViews: newData})
                        } else {
                            set({tabViews: newData, tabViewKey: key})
                        }
                    } else { // 删除当前
                        const index = target.findIndex(item => item.key === key);
                        if (index !== -1) {
                            target.splice(index, 1);
                            if (get().tabViewKey === key) {
                                const next = target[Math.max(0, index - 1)]
                                set({tabViews: target, tabViewKey: next.key as string})
                            } else {
                                set({tabViews: target})
                            }
                        }
                    }
                },
                setTabViewKey: (key) => {
                    set({tabViewKey: key})
                },
                closeTabViewAll: () => {
                    set({tabViewKey: '/home/index', tabViews: [{key: '/home/index', title: '首页', closeIcon: true}]})
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
