学习笔记
Week02 作业题目：

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

参照毛老师讲解的内容。Dao层产生的error应该在Dao层Wrap并上抛，在顶层统一打印log。
同时Dao层应该屏蔽掉底层将业务与底层DB服务解耦,
底层产生的error转化为Dao层的sentinel error。
