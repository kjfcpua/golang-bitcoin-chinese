
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有（c）2013-2014 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

/*
包区块链实现比特币区块处理和链选择规则。

比特币区块处理和链选择规则是一个整体，相当
可能是比特币中最重要的部分。不幸的是，在
这篇文章，这些规则也大部分是没有文件的，必须是
从比特币源代码确定。比特币的核心是
分布式共识：哪些块有效，哪些块将构成
最终确定接受的主区块链（公共分类账）
因此，完全验证节点的一致性非常重要
所有规则。

在较高的层次上，此包支持将新块插入
根据上述规则的区块链。它包括
拒绝重复块、确保块和
事务遵循所有规则、孤立处理和最佳链选择
通过重组。

因为这个套餐不涉及其他比特币的细节，比如网络
通信或钱包，它提供一个通知系统
调用者对某些事件的反应具有高度的灵活性



比特币链处理概述


一系列验证规则。以下列表是
这些规则提供了一些关于引擎盖下发生的事情的直觉，但是
决不是详尽的：

 -拒绝重复块
 -对块及其事务执行一系列健全性检查，例如
   核实工作证明、时间戳、交易数量和特征，
   事务量、脚本复杂性和merkle根计算
 -将块与预先确定的检查点比较预期时间戳
   以及从检查点开始的时间所带来的困难
 -将最近的孤立块保存有限的时间，以防其父块
   块可用
 -如果块作为处理的其余部分是孤立块，则停止处理
   取决于区块链中区块的位置
 -根据块的位置执行一系列更彻底的检查。
   在区块链内，如验证区块困难，坚持
   难以重定目标规则，时间戳位于最后一个的中间值之后
   
   块版本与以前的块一致
 -确定块如何装入链条并执行不同的操作
   因此，为了确保任何难度较大的侧链
   比主链成为新的主链
 
   
   主链），对块的事务执行进一步检查，例如
   
   
   交易价值
 -运行事务脚本以验证是否允许花费者
   硬币
 -将块插入块数据库

错误

此包返回的错误可能是底层提供的原始错误
调用或类型为blockback.ruleerror。这允许呼叫者区分
在数据库错误等意外错误与规则错误之间
通过类型断言违反。此外，调用者可以通过编程
通过检查
类型断言的区块链.ruleError。

比特币改进建议

此包包括以下BIP概述的规范更改：

  bip0016（https://en.bitcoin.it/wiki/bip_0016）
  bip0030（https://en.bitcoin.it/wiki/bip0030）
  bip0034（https://en.bitcoin.it/wiki/bip0034）
**/

package blockchain
