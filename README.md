# Tmp File Cleaner

一時的にファイルを置くフォルダを作って、そのまま放置。それが積み重なって大量のファイルが放置されて、目的のファイルを探すのに時間がかかったり、「このファイルは消して大丈夫？」となったりしていませんか？
Tmp File Cleanerを使って、定期的に一時ファイルを削除してこういった悩みを解消しましょう。

## 使い方

指定されたディレクトリ内のファイルを全て削除します。
ディレクトリの指定はconfig.yamlに記載します。

### config.yaml
YAML形式で記載します。

target : 削除対象の情報を記載します
    - folders : 削除対象のディレクトリパス(絶対パス)を配列で記載します。

``` config.yaml
target: 
    folders: 
        - /path/to/the/target/folder1
        - /path/to/the/target/folder2
```

### 実行オプション

```
    --file, -f:  YAMLファイルを指定します。未指定時は実行ディレクトリ階層内の"config.yaml"を使用します。
```
