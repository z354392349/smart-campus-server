## 数据库文件过大回复失败

1.show global variables like '%max_allowed_packet%'; 查看值  
2.修改 my.ini max_allowed_packet=150M  
3.重启  
4.查看值 是否修改成功

## SQL 多行整理为 1 行

\s+ 替换为
