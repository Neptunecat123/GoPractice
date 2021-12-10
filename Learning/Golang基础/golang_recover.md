# recover函數

 `recover`仅在`defer`函数中有效，在正常的执行过程中，调用`recover`会返回`nil`，无其他效果。如果goroutine陷入恐慌，调用`recover`可以捕获到panic的错误，同时允许后面代码继续运行。
