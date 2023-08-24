import type { IThemes, TItem } from "@/style/type"

const dark = document.querySelector("#dark");

const themes: IThemes = {
    default: {
        backgroundColor: "#E9E9E9",
        backgroundColor2: "#BABABA",
        textColor: "#363636",
    },
    dark: {
        backgroundColor: "#363636",
        backgroundColor2: "#4F4F4F",
        textColor: "aliceblue",
    }
}

/**
 * 修改主题
 * @param target 主题变量
 */
const changeStyle = (target: TItem) => {
    for (const key in target) {
        document
            .getElementsByTagName('body')[0]
            .style.setProperty(`--${key}`, target[key])
    }
}

/**
 * 设置黑暗主题
 * @param theme 主题
 * @returns undefined
 */
const changeDark = (theme:string) =>{
    if(dark === null) return
    if(theme === "dark"){
        dark.className = "dark"
    }else{
        dark.className =""
    }
}

/**
 * 改变主题
 * @param flag 是否初始化
 * @param themeName 主题名称
 */
export const setTheme = (flag: boolean = true, themeName: string = "default") => {
    let theme = localStorage.getItem("theme")
    if (flag && theme) {
        changeStyle(themes[theme])
        changeDark(theme)
    } else {
        localStorage.setItem('theme', themeName)
        changeStyle(themes[themeName])
        changeDark(themeName)
    }
}