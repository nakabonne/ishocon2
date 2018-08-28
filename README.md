# ishocon2

## 作業ログ

11:00 競技開始時 13000

12:00 candidatesにindex貼った 19224

13:00 暫定的にvote数クエリ変更 19978
https://github.com/nakabonne/ishocon2/commit/81c8c72eed5dd6e123d270e500dd2880d83394db

13:20 cssキャッシュしたらなんか下がった 12000

13:40 party毎のvote数をインメモリキャッシュした 20888
https://github.com/nakabonne/ishocon2/commit/600bcd965015d6b1a6c293584e9b8b06877a3147

14:11 投票結果をgoプロセスインメモリキャッシュ 26288
https://github.com/nakabonne/ishocon2/commit/c603f56443aa8b365d825ec85c2c06db58a934ef

14:22 全てのcandidateをgoインメモリキャッシュ 26990
https://github.com/nakabonne/ishocon2/commit/0f5b231c0ed2a154183d8403a96c3776a7f140ab

15:24 voteをバルクインサートに変更 27277
https://github.com/nakabonne/ishocon2/commit/88a213f3c847b294bff59cc226d6709e99372519

18:00userをインメモリに置こうとしたが値の整合性保てず失敗
