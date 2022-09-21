# fileCopier
## 环境
Windows
## 功能
出于个人需求，将文件夹备份到同一台设备的其他内存分区中。
如果有一个经常修改的文件夹`note`被放置于桌面，在每次修改完后运行这个程序，可以把修改后的文件夹的状态同步到D盘，E盘和移动硬盘。处于安全和完备考虑，有以下设定。
## 特点
通过`文件名` : `修改时间`，用于判断那些文件被修改过。
1. 如果`note`文件夹中的一个文件，它曾被备份到了D盘`note2`文件夹；在修改后，这个文件不在`note`文件夹中，那么`note2`会把这个文件挪入`note2`里面一个叫`fileCopier_TrashBin`的子文件夹。
2. 如果`note`中，出现了`note2`没有的文件，它将会被直接备份。
3. 如果`note`中的一个文件，他的修改时间发生了变化，他将被更新到`note2`中。
## 使用
### 添加目标文件夹
请注意`fileCopier.json`文件。
如果想把桌面`note`文件夹备份到D盘、E盘和移动硬盘（G:)，则在Folder中添加：
```json
{
    "source": "C:\\Users\\15129\\Desktop\\note",
    "destination": ["D:\\note", "E:\\note", "G:\\note"]
}
```
最后得到下面的`fileCopier.json`文件。如果有需要，我们可以到`D:\\note`、`E:\\note`、`G:\\note`中去查看备份的文件。
``` json
{
    "Folders": [
        {
            "source": "C:\\Users\\15129\\Desktop\\note",
            "destination": ["D:\\note", "E:\\note", "G:\\note"]
        }
    ],
    "auto": false
}
```
还想把D盘的`diary`文件夹备份到E盘，`fileCopier.json`看起来是这个样子：
``` json
{
    "Folders": [
        {
            "source": "C:\\Users\\15129\\Desktop\\note",
            "destination": ["D:\\note", "E:\\note", "G:\\note"]
        },
        {
            "source": "D:\\diary",
            "destination": ["E:\\diary"]
        }
    ],
    "auto": false
}
```
每次运行程序，都会检查`fileCopier.json`中所有需要备份的文件夹并完成备份。
### 自动运行或手动运行
1. 把`fileCopier.json`中的`auto`设为`true`，把程序添加到[开机自启动](https://support.microsoft.com/zh-cn/windows/%E5%9C%A8-windows-10-%E4%B8%AD%E6%B7%BB%E5%8A%A0%E5%9C%A8%E5%90%AF%E5%8A%A8%E6%97%B6%E8%87%AA%E5%8A%A8%E8%BF%90%E8%A1%8C%E7%9A%84%E5%BA%94%E7%94%A8-150da165-dcd9-7230-517b-cf3c295d89dd)。
2. 或者设为`false`，每次修改后手动运行。