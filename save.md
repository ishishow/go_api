Skip to content
Search or jump to…

Pull requests
Issues
Marketplace
Explore

@ishishow
CyberAgentHack
/
web-speed-hackathon-online
13
173
46
Code
Issues
Pull requests
Actions
Projects
Wiki
Security
Insights
Web Speed Hackathon Online 出題のねらいと解説
nodaguti edited this page on 22 Jul 2020 · 1 revision
皆さま競技お疲れ様でした！

今回のお題アプリケーション「Amida Blog: あみぶろ」は徹底的なデチューニングを施した結果 Lighthouse で安定して 0 点を叩き出す代物でした． あみぶろにはどこにどのような改善ポイントがあったのかについて解説します．

また「どうすれば改善ポイントに気がつくことができたのか」についても，適宜各項目で触れるほか一般的なテクニックを記事の最後で説明します．

※ 各項目は主にジャンル別に配列されており，必ずしも有効な施策順に並んでいるわけではないことにご注意ください

はじめに: パフォーマンスメトリクスについて
ビルド
NODE_ENV
webpack
Source Maps
Mode
Chunk Splitting (ビルド編)
babel-loader
url-loader
Babel
@babel/preset-env
@babel/plugin-transform-modules-commonjs
PostCSS
Source Maps
postcss-custom-properties
postcss-calc
Minify
フロントエンド
Render-Blocking Resources の削減
Resource Prioritisation
Web Fonts
不必要な CSS の削除
Chunk Splitting (フロントエンド編)
依存パッケージの最適化
Polyfills
jQuery
ImmutableJS
lodash
moment-timezone
bluebird & race-timeout
Axios
react-helmet
React & React DOM
処理の並列化
処理のメモ化
処理の遅延実行
ロジックの簡略化
画像の最適化
Lazy Loading
適切なファイルフォーマットの選択
ファイルサイズ・画質・解像度の最適化
余談: ランタイムのパフォーマンス向上
Service Worker
React の Reconciliation 抑制
バックエンド
babel-node を使用しない
不要な処理の削除
Compression
静的ファイルの配信を最適化
キャッシュの設定
サーバーロケーション・サーバースペックの選定
HTTP/2 による配信
Server-Side Rendering (SSR)
余談: サポートブラウザを広げる必要がある場合には
Differential Serving
WebP と picture 要素
gzip と Brotli
パフォーマンスボトルネックを探すには
さらに学びたい方へ
はじめに: パフォーマンスメトリクスについて
具体的な改善方法の紹介に入る前に，まずは採点に使用していたパフォーマンスメトリクスについて簡単に説明します．

「ページの読み込みが速い」というのは，具体的にいつからいつまでにかかる時間が短いことをいうのでしょうか？ Web ページで JS が占める割合が大きくなる中で，従来の onLoad や onDOMContentLoaded ではユーザーの感じる「読み込み速度」との乖離が大きくなってしまい，2016 年ごろから様々なメトリクス（指標）が考案されてきました．

採点に使用した Lighthouse v6 では以下の指標をもとに得点を計算しています．

First Contentful Paint
Speed Index
Largest Contentful Paint
Time to Interactive
Total Blocking Time
Cumulative Layout Shift
それぞれのメトリクスの詳しい定義についてはリンク先を参照してください．

これらメトリクスは不変のものではなく，今でもよりユーザーの体感に合致した指標は何か？と試行錯誤が続けられています．例えば，Largest Contentful Paint, Total Blocking Time などは 2020 年になって提案された新しい指標です．

指標の種類や測定の方法，ここに挙げたもの以外の指標については User-centric performance metrics に詳しくまとまっていますのでぜひ一読をおすすめします．

また，数多くある指標の中から特に Web サイトのユーザー体験に影響があり重要なものとして Web Vitals を Google が提案しています．2020年時点では以下の3つの指標が挙げられており，すべての Web サイトが計測するべきだと位置付け，目標値も公開しています．

Largest Contentful Paint
First Input Delay
Cumulative Layout Shift
それでは，チューニングで改善させるべき数値がわかったところで，さっそく具体的なチューニングポイントを見ていきましょう．

ビルド
NODE_ENV
NPM scripts の build コマンドは

"build:webpack": "cross-env NODE_ENV=development webpack --config webpack.config.js",
となっており， NODE_ENV=development が常に適用されるようになっていました． React などライブラリによっては NODE_ENV の値をもとに開発モード・プロダクションモードを切り替えるものがあります ので， NODE_ENV=production となるよう変更することで改善できます．

- "build:webpack": "cross-env NODE_ENV=development webpack --config webpack.config.js",
+ "build:webpack": "cross-env NODE_ENV=production webpack --config webpack.config.js",
webpack
最近では Parcel などの zero-configuration なバンドラーや Next.js, Gatsby といった最適化を裏で自動的に行ってくれる統合フレームワークが登場してきていますが，生の webpack も webpack.config.js によって bundle の分け方などを細かく設定できることから根強い人気があります．

今回の競技では，ややもするとツールによって隠されてしまいがちな webpack の最適な設定についてどこまで把握できているかを題材にしました．

Source Maps
DevTools でトランスパイル前のコードを参照できるようにする技術として source map がありますが，あみぶろでは

devtool: 'inline-source-map',
という設定になっています．

これは最も詳細な情報が得られるサイズの大きい source map をトランスパイル後のファイルに inline で埋め込むというもので， Devtool | webpack に書かれているとおりファイルサイズが巨大化するため不適切です．

プロダクションモードでは off になるよう変更するか，inline に埋め込まない別の形式に変更することで改善できます．

- devtool: 'inline-source-map',
+ devtool: NODE_ENV === 'production' ? false : 'inline-source-map',
Mode
webpack は v4 から Mode という機能が導入され， mode を切り替えるだけで development/production 用に最適化された設定に変えてくれるようになりました． これは NODE_ENV に連動しているわけではないので，明示的に有効にする必要があります．

あみぶろでは mode による自動設定を無効化する none に設定されていました． production に設定することで，変数名を短くしたり不要なコードを削除したりする TerserPlugin, モジュールの結合を最適化する ModuleConcatenationPlugin などが有効になります．なお，各モードで変更される設定の詳細については上記公式ドキュメントを参照してください．

- mode: 'none',
+ mode: 'production', // or mode: process.env.NODE_ENV,
Chunk Splitting (ビルド編)
Chunk Splitting とは，通常全て一つにまとめられてしまう JS のコードを適宜分割することで各ページごとに読み込まれる JS を最適化することです．これにより，たとえばトップページでは必要ない JS を遅延読み込みするなどの改善が可能になります．

あみぶろでは Chunk Splitting は無効化されているため全てのコードが一つの巨大な main.bundle.js にまとめられるようになっていました． mode: 'production' を有効にすることで webpack がある程度の最適化 を行ってくれるようになります．

より最適化したい場合には個別に optimization.splitChunks の設定を変更することになります．その際にはたとえば以下の Next.js の最適化方法などが参考になります．

Blog - Next.js 9.2 > Improved Code-Splitting Strategy | Next.js
Improved Next.js and Gatsby page load performance with granular chunking
フロントエンドのコードを変更することでさらに chunk splitting を進める方法は後述します．

babel-loader
あみぶろの

{
  test: /\.m?jsx?$/,
  use: {
    loader: 'babel-loader',
  },
},
という設定には exclude が設定されていません．

exclude は loader の除外設定を行うオプションですので，これがないと node_modules の中にあるパッケージも全てこの loader の対象になって全て再トランスパイルされることになってしまいます．そのため，以下のように変更して改善します．

  {
    test: /\.m?jsx?$/,
+   exclude: /node_modules/,
    use: {
      loader: 'babel-loader',
    },
  },
url-loader
url-loader は import されたファイルを Base64 に変換してファイル内に埋め込む loader です．今回は画像の読み込みに使用していますが，画像の最適化が行われない点や Base64 は平均してサイズが元の 130% へ増加する点を考慮すると不適切と言えます．

ひとまず file-loader を使用することで JS ファイルへの画像の埋め込みは避けられますが，より最適化するためには Imagemin など自動的にファイルを圧縮するツールを使用するとよいでしょう (Webpack 経由であるいは CLI で直接使うことができます）．

画像の最適化についてはフロントエンドの章でより詳しく解説します．

Babel
あみぶろの .babelrc は以下のようなシンプルなものですが，ここにも問題点があります．

{
  "presets": ["@babel/preset-env", "@babel/preset-react"],
  "plugins": ["@babel/plugin-transform-modules-commonjs"]
}
@babel/preset-env
@babel/preset-env は targets を設定せず，プロジェクトのどこにも browserslist の設定が書かれていない場合には，以下の説明のとおり ES2015+ のコードを全てトランスパイルします．

if no targets are specified, @babel/preset-env will transform all ECMAScript 2015+ code by default.

今回のターゲットブラウザは Chrome 最新版だけですので， last 1 Chrome major version と .browserslistrc などに指定すれば Chrome 最新版に必要な分だけのトランスパイルが行われるようになり，出力コードが小さくなります．

@babel/plugin-transform-modules-commonjs
webpack は tree shaking という，モジュールから実際に import されたコードだけを残してそれ以外を削除することのできる機能を搭載しています．しかし，この機能を動作させるためには webpack がコードの最適化を行う段階で ESModules 形式 (import/export 文を使った形) になっていなければいけません．言い換えると，Babel で ESModules が変換されてしまっていると tree shaking がうまく動きません．

@babel/preset-env の modules はデフォルトで auto なので問題ないのですが，この @babel/plugin-transform-modules-commonjs があることによって全て CommonJs 形式に置き換わってしまう設定になっています．

そのため，このプラグインは削除する必要があります．

  {
    "presets": ["@babel/preset-env", "@babel/preset-react"],
-   "plugins": ["@babel/plugin-transform-modules-commonjs"]
  }
PostCSS
Source Maps
PostCSS の source map を有効にすると， デフォルトで inline オプションが有効になってしまいます．そのため， inline を明示的に false にするか，source map そのものを無効化することで改善できます．

-  map: true,
+  map: false, // or map: { inline: false },
postcss-custom-properties
postcss-custom-properties は CSS Custom Properties (いわゆる CSS Variables のような機能) をトランスパイルするためのプラグインです．

Chrome 最新版は Custom Properties をネイティブでサポートしているため，ターゲットブラウザを考えるとこのプラグインは削除することも可能ですが，変数が使われている箇所を値に置き換えるためコードの最小化という観点から考えると実は有用なプラグインです．

しかし，デフォルトの設定では preserve(https://github.com/postcss/postcss-custom-properties#preserve) というオプションが有効になっています．これは，以下のように置換前と置換後の両方を残しておくという設定で，コード量がむしろ倍に増えてしまいます．

:root {
  --color: red;
}

h1 {
  color: var(--color);
}

/* becomes */

:root {
  --color: red;
}

h1 {
  color: red;
  color: var(--color);
}
そこで preserve: false を設定することで置換後のプロパティだけが残るようにします．

-  customProperties(),
+  customProperties({ preserve: false }),
:root {
  --color: red;
}

h1 {
  color: var(--color);
}

/* becomes */

h1 {
  color: red;
}
postcss-calc
postcss-calc は calc() が使われている部分をコンパイル時に事前計算して置き換えることによりコード量を削減するプラグインです．今回のアプリケーションでも随所で calc() が使われているため有効です．

+  const calc = require('postcss-calc');
...
+    calc(),
Minify
例に漏れず CSS も minify されていません． cssnano が PostCSS 用の minifier として有名です．

+  const cssnano = require('cssnano');
...
+    cssnano(),
フロントエンド
Render-Blocking Resources の削減
webpack の html-webpack-plugin によって生成される HTML は，script を同期的に読み込むようになっていました．

<script
  type="text/javascript"
  src="main.bundle.js"
></script>
この場合，ブラウザは <script> に遭遇した段階で HTML のパースを停止し，JS の取得・実行を行ってから残りの HTML を処理します．こうしたレンダリングを阻害するリソースは Render-Blocking Resources と呼ばれ，それらを可能な限り減らすことでレンダリングの高速化につながります．

<script> の場合には defer 属性と async 属性という似たような指定が存在しますが， この二つは挙動が少し異なります． HTML の仕様書に書かれている図がわかりやすいのですが， async を指定した場合は HTML のパースと JS の取得を並行で行うものの，取得が完了した段階で HTML のパースを止めて即座に JS を実行します． defer では並行に JS の取得を行ったのち，HTML のパースが終了した段階(つまり DOMContentLoaded が発火する段階)になってから JS を実行します．

HTML のパースが中断する可能性のある async より defer の方が描画が早く始まり，かつ実行されるタイミングも安定しているため好まれます．

  <script
    type="text/javascript"
    src="main.bundle.js"
+   defer
  ></script>
この指定により app.js で onload を待っている箇所も削除することができます．

-  window.onload = () => {
     init();
-  };
一方で CSS はやや複雑です．ブラウザは HTML のツリーと CSS のスタイル情報がないと描画を始められないため，できる限り初回に読み込まれる CSS のサイズは小さい方がよいのですが，全て遅延させてしまうとページ読み込み直後に全く CSS が適用されていないページが表示されてしまうことになります．

必要最小限の CSS のみを読み込むようにする技術として，above-the-fold (ページ全体のうち，読み込み直後に画面内に入っている部分) の表示に必要な CSS だけを抜き出し，それ以外の CSS を遅延読み込みさせる方法があります．以下の記事ではそうした技術が紹介されています．

https://developers.google.com/web/fundamentals/performance/critical-rendering-path/render-blocking-css
https://web.dev/defer-non-critical-css/
https://web.dev/extract-critical-css/
Resource Prioritisation
HTML にはもう一つ改善点が存在します．それは resource prioritisation (リソースの優先度付け) の設定です．

ブラウザがページを読み込む際には，リソースの種類や読み込みの方法などで自動的に読み込みの優先度を決定して処理をしています．たとえば，Chrome の 2016 年時点での優先度付けの方法は Preload, Prefetch And Priorities in Chrome - reloading - Medium にて解説されているほか，DevTools では個々のリソースについて優先度の判断の結果を見ることができます:

Priorities in Chrome Developer Tools

なお，上記はあくまで Chrome の処理内容であって，標準化されたものではないため，Firefox や Safari など異なるブラウザでは異なる優先度付けをすることに注意が必要です．

これらブラウザによるデフォルトの優先度付けに対して，開発者側からカスタムで優先度付けのヒントを提供することができます．その仕組みが Resource Hints です．これは <link> タグを通じて特定のリソースの優先度をブラウザに示唆するものです． <script> や <img>，JS によって事前読み込みを行う仕組みとは異なり，直接ブラウザのリソース読み込み優先順に働きかけることができますが，あくまで "Hints" なので実際に指定されたとおりに処理されることが保証されているわけではありません．

ブラウザへ示唆する内容によっていくつか種類があります:

dns-prefetch: 指定したドメインに配置してあるリソースを読み込み中のページ内で使用するので，DNS ルックアップを事前に行っておくとよい旨をブラウザに示唆します．
preconnect: 指定したドメインに配置してあるリソースを読み込み中のページ内で使用するので，コネクションを事前に確立しておくとよい旨をブラウザに示唆します．これには DNS ルックアップ, TCP ハンドシェイク，TLS ネゴシエーションが含まれます．
prefetch: 指定したリソースを読み込み中のページから遷移した先で使用するので，事前に読み込んでおくとよい旨をブラウザに示唆します． prefetch が指定されたリソースは読み込み中のページでは使われないはずなので，一般にブラウザは優先度をデフォルトよりも低く設定します．
prerender: 指定したページへ読み込み中のページから遷移する可能性が高いので，事前に読み込んでレンダリングしておくとよい旨をブラウザに示唆します．ただし，レンダリングに必要な帯域消費やメモリ消費量と比べて実際に得られるパフォーマンス改善度合いが高くないので，Chrome では実際にレンダリングする代わりにページに使われているリソースを全て prefetch する NoState Prefetch という挙動をします．
ブラウザによってサポート状況に違いがあることから dns-prefetch と preconnect は同時に指定するとよいと言われています．

また，Resource Hints とは独立した仕様ながら類似の用途に使われるものとして Preload があります． preload が指定されたリソースは，ブラウザに対して読み込み中のページで使用されるので優先的に読み込みを開始するべきであることを指示します．これは Resource Hints と異なり示唆ではなく指示であるため，ブラウザは必ず優先順を変更し，最優先で読み込みを開始します．

これらの指定に関するユースケースについてさらに詳しく知りたい場合には

Preloading content with rel="preload" - HTML: Hypertext Markup Language | MDN
Resource Prioritization – Getting the Browser to Help You
Preload, Prefetch And Priorities in Chrome - reloading - Medium
などを参照してください．

あみぶろでは画像，SNS シェアボタン，Web Fonts などが外部ドメインから読み込まれているので，それらに適切に dns-prefetch や preconnect, preload を設定することで高速化が見込めます．ただし， 闇雲に設定するだけでは却って逆効果になることもありますので，設定にあたっては DevTools をみてリソースの読み込みがどう変化するかを観察しながら行うことをおすすめします．

一方 Priotiry Hints や 103 Early Hints も Resource Hints と似た名前をしているパフォーマンス関連の仕様ですが，内容は異なります．

Priority Hints は Resource Hints を補完する目的で Google により提案されている仕様で， <script>, <img>, <link>, fetch() などに importance というパラメータを付与して priority を制御できるようにするものです． Chrome 70 以降で実装されていますが，Web Platform Incubator Community Group (WICG) によって提出されているにすぎず W3C の標準化トラックに乗っているわけではありません．Chrome 以外のブラウザでの実装や，実装の計画も現状ありません (例えば Mozilla の反応)．Chrome による初期段階の調査では特に効果が見られなかったという報告もあります．しかしながら，もし設定しようとする場合には Resource Hints や Preload の場合と同じく DevTools の resource waterfall を見ながら調整した方がよいでしょう．

103 Early Hints は HTTP の Status Code として新しく提案されているもので，HTTP/2 の Server Push が使えない HTTP/1.1 において，body をサーバーサイドで生成し終わるより前にクライアントに対して事前に読み込みを開始してほしいサブリソースを返すための status です．現在のところ残念ながらサポートしているブラウザはありません (Chrome, Firefox)．

Web Fonts
あみぶろでは以下の @import 文によって Google Fonts 経由で Web Fonts を使用しています．

@import 'https://fonts.googleapis.com/css?family=Baloo+Thambi+2:400,500,600,700,800&display=swap';
@import 'https://fonts.googleapis.com/css?family=M+PLUS+Rounded+1c:400,500,600,700,800&display=swap';
ここには2つの問題があります．

まず一つ目は不要な weight のフォントも読み込んでいるという点です．アプリケーションをよく精査すると，実はこの Web Fonts は font-weight: bold; でしか使われていないことがわかります． bold は 700 に等しいので， 700 だけ読み込めば充分です．

また，この指定はアプリケーションで使われていない文字のグリフも読み込んでしまうという点でも非効率です．Google Fonts では &text= を指定することで特定の文字だけを読み込むようにできる機能があるので，それを使えば最適化できます．

-  @import 'https://fonts.googleapis.com/css?family=Baloo+Thambi+2:400,500,600,700,800&display=swap';
-  @import 'https://fonts.googleapis.com/css?family=M+PLUS+Rounded+1c:400,500,600,700,800&display=swap';
+  @import 'https://fonts.googleapis.com/css?family=Baloo+Thambi+2:700&display=swap&text=Amida%20Blog:';
+  @import 'https://fonts.googleapis.com/css?family=M+PLUS+Rounded+1c:700&display=swap&text=あみぶろアミブロ阿弥';
これらの設定は最終的に読み込まれるフォントファイルだけでなく，フォントを読み込むために Google Fonts から配信される CSS の量も削減することができます．

不必要な CSS の削除
CSS コード量の削減は，コードのダウンロード・パース・実行の全てにかかる時間を短くすることができるため非常に重要です．あみぶろには二つの方法で不要な CSS が紛れ込んでいました．

一つ目は使われていない大量の utility class です． DevTools の Coverage 機能を使用したり，バンドル後の CSS ファイルに書かれているクラス名を検索したりするとわかるのですが，実は utils.css にて import されていた

@import 'suitcss-utils';
@import '@zendeskgarden/css-utilities';
@import '@zendeskgarden/css-buttons';
@import '@zendeskgarden/css-forms';
は全て未使用のものでした．そのため，これはファイルごと削除できます．

  @import './foundation/styles/vars.css';
  @import './foundation/styles/global.css';
  @import './foundation/styles/web_fonts.css';
- @import './foundation/styles/utils.css';
  @import './foundation/components/foundation_components.css';
  @import './domains/domains.css';
  @import './pages/pages.css';
また，個別のコンポーネントにも以下のような水増しされた不要な CSS が数多く書かれていました:

/* EntryView.css */

.entry-EntryView__figcaption {
  font-size: var(--font-size-s);
  margin-top: calc(var(--space) * 2);
  text-align: center;
}

/* こちらは使われていない */
.entry-EntryView__caption {
  font-size: var(--font-size-s);
  margin-top: calc(var(--space) * 2);
  text-align: center;
}
個々のファイルに書かれているセレクターのうち未使用のものを見つけ出す方法はいくつかあります．いくつか例を挙げると

手動: DevTools の Coverage 機能
静的な HTML/CSS の場合: UnCSS や DropCSS などのツール
CSS-in-JS の場合: minifier や tree shaking により自動的に未使用 CSS を削除できることが多い
などです．

今回のようなコンポーネントごとに JSX と CSS があるような場合には，拙作の stylelint-no-unused-selectors が使えます．これは stylelint のプラグインとして動作し，未使用の CSS を警告してくれます．

Chunk Splitting (フロントエンド編)
Webpack の設定による自動的な Code Splitting に加え，コード上で開発者が明示的に chunk を分割するよう指定することもできます．それが Dynamic Imports を使ったコードの分割です．

SPA ではページの route 単位で分割することがよく行われます． https://example.com/a にアクセスした際に https://example.com/b に必要な JS まで取得するのを止めるイメージです．

Dynamic Imports は Promise を返すため，react-router を使っているあみぶろの場合には React 上でどうにかして Promise をハンドリングする必要があります．よく使われるのは loadable-components というライブラリです．最近 React に入った React.lazy と React.Suspence を使う方法 もありますが，これらは記事執筆時点ではまだ experimental であるため，プロダクション環境で用いるのは避けるのが賢明です．

ちなみにパフォーマンス改善からは外れますが，loadable-components や React.Suspence を使う場合には fallback として読み込みしている間に表示する内容を指定する必要があり，ページ遷移時に一瞬白い画面が出てしまう可能性があります．UX の観点から読み込み中は遷移前の画面を表示するなど routing 挙動をより細かく制御したい場合には universal-router のようなライブラリの使用を検討するとよいでしょう．

依存パッケージの最適化
外部パッケージは便利なものですが，盲目的にあれもこれもと使っているとバンドルサイズが肥大化してしまいます．サイズ，メンテナンス性などを含めて総合的に使用するパッケージを選定することが重要です．

既存のプロジェクトにどのようなパッケージが含まれているのかを知るには webpack-bundle-analyzer や npm ls が有用です．前者は各バンドルにどのパッケージ・ファイルが含まれているのかをタイル形式でグラフィカルに表示してくれます．サイズの大きいファイルほどタイルのサイズが大きくなるため，削減の効果が高いファイルが何なのかを判断できます．後者はパッケージの依存関係を判断するのに使えます．webpack-bundle-analyzer ではしばしば package.json に直接記載されていないパッケージを目にすることがあります．そうしたパッケージがなぜ含まれているのかを npm ls を使うことで効率よく調べられます．

一方で新しいパッケージを追加しようとしている時やサイズの大きいパッケージの代替物を探している時には Bundlephobia が便利です．パッケージ名を入力するとそのパッケージの minifed および minified & gzipped のサイズを表示してくれます．

「サイズだけ表示されてもそれが果たして大きいのか適切なのか分からない」という場合には，Google の提唱する JS のコードサイズは 170KB に納めるべきという指針が目安になるでしょう．

あみぶろでは意図的にサイズが大きかったり本来不要であったりするパッケージを含んでいます．あみぶろで使われているパッケージのうち，削除・置き換えが可能なものの例を以下に挙げます．サイズは執筆時点での最新版の minified & gzipped の値を表します．

Polyfills
core-js のフルバージョンと regenerator-runtime が入っており，それぞれ 45.5 KB と 2.3 KB を占めます．

モダンなブラウザでは大半の polyfill は不要なので，一律で全てを読み込むようにするのはかなり非効率です．

今回のようにターゲットブラウザがいわゆるモダンブラウザだけに限定されている場合には， @babel/preset-env の useBuiltIns と corejs オプションで必要なものだけをバンドルするように設定するのが簡単でしょう．

一方 IE や古い iOS Safari などを含む場合には，ブラウザごとに自動的に必要な polyfill だけを配信してくれる https://polyfill.io が便利です．ただし polyfill.io は SLA を提供していないため，アプリケーションの起動に必須な polyfill が配信されなくなる可能性があることでサイト全体の SLA の定義や SLO の達成が困難になることも考えられます．また，単純に third-party への通信が発生することによるオーバーヘッドを厭忌することもあるでしょう．そうした場合には，polyfill.io が裏で使っている polyfill-library を使うことで比較的簡単にバックエンド (BFF) にて自前で同様の仕組みを構築できます．

jQuery
30.4KB を占めます．React DOM のマウントポイントを取得する部分と，SNS シェアボタンのスクリプトをページに inject する部分であえて jQuery を使用していました．

const root = $('#root').get()[0];
const script$ = $(
  `<script crossorigin="anonymous" src=${FACEBOOK_SDK}></script>`,
).appendTo('body');

return () => {
  script$.remove();
};
これらはもちろん通常の DOM 操作に置き換えることが可能です．

ImmutableJS
17.2KB を占めます．あみぶろでは API から取ってきた JSON を単に Map や List でラップして store に保存し，コンポーネントで toJS() するだけというほとんど意味のない使い方をしています．

そのため，単純に object や array を spread syntax などを使って immutable に更新するだけで充分で，ImmutableJS は削除できます．

lodash
24.3KB を占めます．

まず，lodash をそのまま使用するのはパフォーマンス観点からは悪手です．ESModules 形式で export されていないため，tree shaking が動作せず一つのメソッドを使用するだけで全ての関数がバンドルされてしまうからです．Tree shaking に対応した lodash-es を使うようにするか，babel-plugin-lodash, lodash-webpack-plugin, babel-plugin-import, babel-transform-imports などを使って tree shaking ができる形に変換するのがよいでしょう．

しかし実はあみぶろで使っている lodash は map, filter, chunk, take, shuffle くらいしかないためそもそも自前で書いてもさほど大変ではなく，lodash を完全に消し去ることも可能です．

moment-timezone
94.9KB を占めます．

まず timezone のデータが巨大なことに気がつくと思います．これは Huge file size when using webpack · Issue #356 · moment/moment-timezone に書かれた手法を使って JP 以外のデータを削除することでかなり削減可能です．

しかし，実はあみぶろのコードを調べると timezone の機能を使っている箇所は全くないことがわかります．そのため，素の moment (20.4KB) に置き換え可能です．

さらに突き詰めるとそもそも moment である必要もないので，dayjs(2.8KB)のような軽量ライブラリに置き換えることもできます（相対時刻表示を行う箇所があるので，全て自前に置き換えるのは少し難しいと思われます）．

bluebird & race-timeout
bundle analyzer の結果を見ると bluebird (21.7KB) も目につきます．これは Promise の polyfill と関連するユーティリティ関数を集めたようなライブラリですが，polyfill は別途設定している上 Chrome ではそもそも必要ありません．

このパッケージの削除はやや難しく，あみぶろのコードや package.json を探しても bluebird を直接触っている箇所は見つかりません．つまり，何らかのパッケージがさらに依存している先で使用していることが考えられます．

そのような場合に使えるのが前述した npm ls です．これを実行すると以下が得られます:

$ npm ls bluebird --depth=10
web-speed-hackathon-online@0.0.1 /path/to/web-speed-hackathon-online
└─┬ webpack@4.42.0
  └─┬ terser-webpack-plugin@1.4.3
    └─┬ cacache@12.0.3
      └── bluebird@3.7.2
webpack の依存から入っていることが分かりましたが，クライアントで使っているパッケージで直接依存しているものはないようです．そこで yarn.lock の中を bluebird で検索してみると

race-timeout@^1.0.0:
  version "1.0.0"
  resolved "https://registry.yarnpkg.com/race-timeout/-/race-timeout-1.0.0.tgz#2c20c246662b9748aec1d7b4af4b90406e8f857e"
  integrity sha1-LCDCRmYrl0iuwde0r0uQQG6PhX4=
  dependencies:
    native-or-bluebird "1"
という項目が見つかります．この native-or-bluebird は依存ツリーの中に bluebird が存在するとそれを自動的に使うようにしてしまう厄介な存在です．

そこで race-timeout をどこで使っているのかを調べると，gateway.js にて timeout 処理を実現するために用いていることがわかります．

const requestWithTimeout = timeout(axios.get(path), TIMEOUT);
timeout は axios の設定でも実現できますし， Promise.race() と setTimeout を組み合わせて自前で実装することも簡単です．

これで race-timeout (651B) もろとも bluebird を消し去ることができました．

Axios
Axios はその利便性から広く使われているライブラリですが，4.4KB を消費します．アプリケーションによっては単純な native の fetch() で充分なことも多く，あみぶろでも fetch() に置き換えて Axios は完全に削除できそうです．

もし Axios を完全に捨て去ることが難しいという場合でも，同じインターフェースを提供している軽量な代替である redaxios (884B) への置き換えが可能か検討してみる価値はあります．

react-helmet
サイズは 5.9KB です．JSX を使って宣言的に <head> 内の各要素を書き換えることのできる便利なライブラリですが，あみぶろでは <title> を書き換える用途にしか使っていません．Hooks を使って document.title 経由で操作すれば充分でしょう．

React.useEffect(() => {
  document.title = `${entry.title} - ${blog.nickname} - Amida Blog: あみぶろ`;
}, [entry, blog]);
React & React DOM
あみぶろの根幹をなしている React ですが，React で 2.6KB，React DOM で 35.9KB と軽量とは言い難い大きさをしています．

軽量な代替としては Preact (3.8KB) が有名です．軽量版 React を目指したプロジェクトは数多くありましたが，Preact は今でも活発にメンテナンスが続けられ，React の最新 API に追従しています．

React との差異は Differences to React | Preact にまとめられています．あみぶろのユースケースであれば充分に置き換え可能です．

処理の並列化
あみぶろの各ページを表示するためにデータを取得する部分がありますが，それは全てシーケンシャルな挙動になっています．例えばブログのエントリーを表示する Entry.jsx は以下のとおりです．

await fetchBlog({ dispatch, blogId });
await fetchEntry({ dispatch, blogId, entryId });
await fetchCommentList({ dispatch, blogId, entryId });
これらのデータには特に依存関係がないため， Promise.all() を使って並列に取得することが可能です．

await Promise.all([
  fetchBlog({ dispatch, blogId }),
  fetchEntry({ dispatch, blogId, entryId }),
  fetchCommentList({ dispatch, blogId, entryId }),
]);
処理のメモ化
jQuery の時代から言われていることではありますが，処理のメモ化 (memoise, 処理の結果を保存しておいて再利用すること) はパフォーマンス改善において重要です．

あみぶろでは例えば，時刻表示の部分が挙げられます:

<time
  dateTime={moment(comment.posted_at).toISOString(true)}
  title={moment(comment.posted_at).toISOString(true)}
>
  {moment(comment.posted_at).fromNow()}
</time>
moment(comment.posted_at) という部分を変数に入れておけば使い回すことができますし，そもそも dateTime と title 属性で同じ値を二回計算しています．

なお，本筋からは外れますが dateTime 自体を削除することはこの例ではできません． <time> の dateTime 属性を省略できるのは子要素が特定の形式で書かれているときに限られ， 8 days ago のような表記の場合には dateTime が必須となります．同様に <time>10:23</time> のようなフォーマットで動画の長さ (duration) をマークアップしている例も巷で散見されますが，duration は仕様で定義されている PT10M23S か 10m 23s というフォーマットで書かれる必要があり， 10:23 では時刻と解釈されてしまいます．基本的には <time> には dateTime を指定するものだと認識しておいた方がよさそうです．

React に特有のメモ化テクニックとしては React.useMemo が挙げられます．これは第二引数に指定した deps が変更された時のみ指定した関数を実行するというメモ化のための Hooks です．あみぶろではそこまで処理に多大な時間のかかる処理をしていないため効果は限られますが，原理的には上の moment の例でも moment(comment.posted_at) を毎回実行するコストより comment.posted_at の文字列比較をするコストの方が低いため， useMemo による恩恵を受けられるはずです．

処理の遅延実行
全てのデータを取得し終わるまで何も表示せず白い画面が続いてしまうと，ユーザーの体感的な表示速度は低下してしまいます．そうした「ユーザー体験に基づいた表示速度 (Perceived load speed)」を計測しようとするメトリクスである First Meaningful Paint や Largest Contentful Paint もその状態では向上が見込めないでしょう．

あみぶろでは以下のコードに示す hasFetchFinished という state の存在により，データが取得されるまでヘッダーを除いて画面は白いままになってしまいます．

const [hasFetchFinished, setHasFetchFinished] = useState(false);

useEffect(() => {
  setHasFetchFinished(false);

  (async () => {
    try {
      await fetchBlog({ dispatch, blogId });
      await fetchEntry({ dispatch, blogId, entryId });
      await fetchCommentList({ dispatch, blogId, entryId });
    } catch {
      await renderNotFound({ dispatch });
    }

    setHasFetchFinished(true);
  })();
}, [dispatch, blogId, entryId]);

if (!hasFetchFinished) {
  return (
    <Helmet>
      <title>Amida Blog: あみぶろ</title>
    </Helmet>
  );
}
こうした挙動を緩和する一つの手段は，重要ではない要素を遅延取得・遅延表示することです．大半のスクリーンサイズで First View に入らないと考えられるものは Intersection Observer API を利用して要素が実際に画面内に入ろうとしたときに取得するようにできます．仮に First View に入っていたとしても，例えば YouTube が動画プレイヤーを最優先で表示してその他の要素は遅延させているように，ユーザーにとって重要度の低い要素は遅延させることを検討できます．あみぶろであればコメント欄や SNS シェアボタンなどは遅延させてもさほど大きな体験上の劣化を招かないでしょう（レギュレーションでは遅延読み込みを許容するため ページをスクロールしたときに得られる情報に差異がない という表現をしていました）．エントランスのブログ一覧やブログトップの記事一覧など縦に長くなりがちなページでは，Virtual Scroller という類似した考え方も参考になります．

遅延させない場合であっても， hasFetchFinished のような仕組みでページを全て隠すのではなく，取得できたデータから随時表示すればユーザーにとっての体感速度は上がります．その際，データが読み込み中の間は Content Placeholder と呼ばれる代わりの要素を表示しておくと親切です．Placeholder により「ここに表示されるデータは読み込み中で，取得後はこういう UI が出る予定です」ということを予告しておくと，いわゆる「ガタン問題」を防ぐことにもつながり，ユーザーを不安にさせたり驚かせたりすることがなくなります．ちなみにパフォーマンスからは逸れますが，最近では要素の位置が急に変化しないことを測るメトリクスとして Cumulative Layout Shift が Google から提案されるなど，「ガタン問題」は UX の観点からも重要視されています．

その他，App Shell というアーキテクチャも参考になるでしょう．

ロジックの簡略化
簡潔なロジックで要件を実装できればコード量の削減につながり，パフォーマンスとともにリーダビリティも向上します．過剰なコードゴルフはリーダビリティを損なう上に minify や gzip 化によって差がほとんどなくなることも多いのであまり褒められたものではありませんが，通常のコーディングの範囲でより簡単な実装方法があるならば採用しない理由はありません．

あみぶろではエントランス画面にあるブログ一覧をグリッド状で表示する部分に改善の余地があります．このリストは _.chunk() と flexbox を使って以下のように実装されています．

export function BlogCardList({ list, columnCount }) {
  const rows = _.chunk(list, columnCount);

  return (
    <div className="blog-list-BlogCardList">
      {_.map(rows, (rowItems, i) => (
        <div key={i} className="blog-list-BlogCardList__row">
          {_.map(rowItems, (item, j) => (
            <div
              key={j}
              className="blog-list-BlogCardList__column"
              style={{ width: `calc(100% / ${columnCount})` }}
            >
              <BlogCard blog={item} />
            </div>
          ))}
        </div>
      ))}
    </div>
  );
}
しかし，これは CSS Grid を使えば JSX/CSS どちらのコードもずっと削減できます．

export function BlogCardList({ list }) {
  return (
    <div className="blog-list-BlogCardList">
      {
        list.map((item, i) => (
          <BlogCard key={i} blog={item} />
        ))
      }
    </div>
  );
}
.blog-list-BlogCardList {
  display: grid;
  grid-gap: calc(var(--space) * 2);
  grid-template-columns: repeat(4, 1fr);
}
(本筋とは関係ないですが， columnCount も常に 4 なので消すことができます．もしカスタマイズ性を残しておきたい場合には style 属性で Custom Properties を指定し，CSS 側で参照するテクニックが使えます．)

画像の最適化
画像はしばしば KB 単位でファイルサイズが増減するため，コードの最適化よりも影響が大きいことがあります．

Lazy Loading
First View の外にある画像の読み込みを遅延させることにより，そもそもページ読み込み時の画像の数を減らす改善です．

最も手っ取り早い方法としては <img> 要素の loading 属性が挙げられます． loading="lazy" を指定するだけで Lazy Loading が実現でき，ブラウザのサポートも広がってきています．しかし，lazyload のストラテジ（どのタイミングで読み込むか）は仕様に書かれていないため，ブラウザの実装依存となっています．実際，Chrome は回線状況などによって変化するものの viewport からかなり離れたところにある画像も読み込みの対象となりますが，Firefox は viewport 内に入ってから読み込む実装となっています (https://mathiasbynens.be/demo/img-loading-lazy で挙動を確認できます)．本来こうした挙動はブラウザに任せることでブラウザベンダが随時更新する最適なものを享受するのが望ましくはあるのですが，今回の競技に限っていうと Lighthouse が計測に Chromium を使っていること，縦幅が短いページが多いことなどからより繊細な lazyload の制御を行った方がよい結果が得られます．

そのように lazyload のストラテジの制御を開発者側で行いたい場合には前述の Intersection Observer API が便利です． scroll イベントの監視を行うような昔ながらのやり方に対してハイパフォーマンスに動作します．

ところで画像の lazyload は広く知られた技術であるためそれを簡便に実現する React コンポーネントライブラリも多く 存在 しますが，中には react-lazy-load-component のように独自の CSS を自動的に追加してしまうコンポーネントもあるのでサイトの既存の CSS とコンフリクトを起こさないか注意が必要です．あみぶろでは画像のほとんどが ProportionalImage というアスペクト比を維持した状態で指定サイズに納めて表示するコンポーネント経由で表示されていましたが，そうした CSS と干渉を起こしてしまう可能性もあります．

適切なファイルフォーマットの選択
ファイルフォーマットを適切に選択することで，同じ画質でもファイルサイズを削減することができます．例えば WebP は JPEG/PNG の同じ画質の画像と比べて 25-35% もファイルサイズが小さいと言われています．また，アニメーション GIF は MPEG や WebM に変換することで 1/10 程度までサイズを減らすことができます．

あみぶろではエントランスの Hero 画像に最適化されていない PNG 画像が，404 ページにはアニメーション GIF が存在したので，それらを WebP や WebM に変換することで改善可能でした．

WebP への変換は Imagemin，WebM への変換は ffmpeg などを使って実行できます．

ファイルサイズ・画質・解像度の最適化
ファイルサイズ，画質，解像度もそれぞれ最適化が可能です．

JPEG や PNG など既存の画像フォーマットであっても，エンコーダーの改良や不要な領域の削除などでファイルサイズを削減できることがあります．まずは Imagemin を実行してみるのがよいでしょう．

画質や解像度についても，実際に表示される画像の大きさによって最適なものを選択することで大幅なファイルサイズ削減につなげられます．あみぶろではブログの Hero 画像，記事中の画像，コメント欄に表示されるユーザーのアバターなどがことごとく数千px x 数千 pxの巨大な画像になっており，1MB を超える画像も少なくありませんでした．目で見て違いがはっきり分かるほど画質や解像度を落としてしまってはもちろんいけませんが，適切な範囲で調整することで軽量化が図れます．

レポジトリ内に含まれる画像はコンパイル時に Imagemin を実行することで最適化できますが，API から返却される画像の最適化はどうすればよいでしょうか．API の結果は時間とともに変わる可能性があるため，事前に最適化処理を走らせることは困難です．

そのような場合には画像用の CDN を利用すると動的な変換が可能になります．Thumbor という OSS の画像 CDN 実装を用いて自前で CDN サーバを構築することもできますし，imgix や Cloudiary のような SaaS を使うこともできます（ちなみに CyberAgent では Hayabusa という内製の CDN が使われることが多いです）．

余談: ランタイムのパフォーマンス向上
今回の競技では残念ながらスコアに反映されないものの，実際にユーザーがサイト内を回遊する時の体験をパフォーマンス観点から向上させたい場合に行うとよい施策を紹介します．

Service Worker
幅広い用途に用いることのできる Service Worker はクライアントサイドのキャッシュ機構も構築できます．Service Worker スクリプトは初訪時にブラウザへ登録されて再訪時から動作を行うため Lighthouse の計測には反映されませんが，静的なファイルをキャッシュすることも動的な通信 (i.e., fetch() の結果) をキャッシュすることもできます (その究極の形がオフライン対応です)．

素の Service Worker はやや扱いにくいのですが，workbox というライブラリを使うとよくあるユースケースを簡単に実現できるようになります．

React の Reconciliation 抑制
React の VirtualDOM が更新されると内部では Reconciliation という処理が走って前回との diff を計算し，変更があれば実際の DOM に反映させます．実 DOM の操作に比べると軽い処理ではあるものの，ブラウザのメインスレッドで実行されることから UI-blocking な処理であり，高頻度に発生するとアプリケーションの操作性 (responsiveness) へ無視できない影響が発生してしまいます．

React DevTools の Profiler で Highlight updates when components render. という設定を有効にすると reconciliation が発生したコンポーネントをハイライトしてくれるため調査に役立ちます．あみぶろでも，例えばエントランスページでは Hero 画像の上にある「あみぶろ・アミブロ・阿弥ブロ」のアニメーションが動くたびに下にあるブログ一覧の一つ一つのカードが全て更新されているということがわかります．

このように React のデフォルトのコンポーネントの挙動では，どこか一箇所で re-render が発生するとその子孫コンポーネントはたとえ props/state が全く変わっていなかったとしても全て reconciliation が走ってしまいます．この reconciliation を抑制するための仕組みが shouldComponentUpdate，React.PureComponent，React.memo です．

これらは props や state が変更されたかどうかを判断し，変更されていないときには reconciliation を発生させないように制御できます．実装の際には意図せずして props/state の参照が更新されないよう注意が必要です．例えば () => {} でコールバック関数を作成したり object や array を render() 内で作成したりしていると，毎回新しい参照が作られることになって「変更があった」と判断され，reconciliation の抑制が行われなくなってしまいます．React.useMemo，React.useCallback，クラスメソッドを constructor で bind() するイディオムなどを用いて適切に参照を維持しなくてはいけません．

また， shouldComponentUpdate などの設定により re-render ごとに毎回比較処理が走ることになるため，高頻度に props/state が更新される場合や比較処理のコストが大きい場合には却ってパフォーマンスが悪化する可能性もあります．導入の際には Profiler を見ながらの慎重な調整が不可欠です．より詳細な情報は

Optimizing Performance > Avoid Reconciliation – React
Use React.memo() wisely
React製のSPAのパフォーマンスチューニング実例 | リクルートテクノロジーズ　メンバーズブログ
などを参照ください．

バックエンド
あみぶろでは lib ディレクトリの中に実装されていた Backend for Frontend (BFF) 層の実装にも改善点がありました．

ちなみにバックエンドで行われた処理の時間計測はしばしばフロントエンドと切り離されてしまうことも多いですが，Server Timing API を実装することで DevTools にてサーバーサイドの情報も同時にみることができるようになります．Fastly のような CDN でも実装するところが出てきています．

babel-node を使用しない
BFF は静的なトランスパイルを行わず，実行時に動的に変換して動作する babel-node を使って動いていました．

"serve": "nodemon --exec babel-node lib/server.js"
web: babel-node lib/server.js
公式ドキュメントに

You should not be using babel-node in production. It is unnecessarily heavy, with high memory usage due to the cache being stored in memory.

とあるとおりこれは避けるべき操作です． build:server を NPM scripts に追加するなどして事前にサーバーサイドもトランスパイルしておき，実行は通常の Node を用いるのがよいでしょう．

不要な処理の削除
実は BFF にはクライアントで使われない id というフィールドを計算する処理が存在しており，これを削除することで応答速度の向上が見込めました．

function createId(n) {
  const c = [];
  const len = n * 1000;
  for (let i = 0; i < len; i++) {
    c.push[i];
  }
  const result = c.sort((a, b) => a - b).join(',');
  return result;
}

const id = createId(Math.floor(Math.random() * this.data.length));
コードのとおりこの計算処理に必要な時間は一定ではなかったため，謎のゆらぎに苦しめられた方もいたかもしれません．

Compression
あみぶろの BFF はリソースを何も圧縮せずに配信していました．Express であれば compression という middleware で簡単に gzip/deflate 圧縮を追加できます．

また，モダンなブラウザでは Brotli というより効率のよい圧縮形式をサポートしています．Brotli は圧縮にやや時間がかかるものの，Akamai の調査では gzip と比べてサイズの点で 15-25% 圧縮効率がよいと報告されています．こちらは compression の fork である shrink-ray という middleware で実現できます．

一方でこうした middleware はアクセスのたびに動的に圧縮を行うのでやや非効率な点も否めません．さらにパフォーマンスを向上させるには事前に静的に圧縮しておくのがよいでしょう．compression-webpack-plugin や brotli-webpack-plugin を使うことでコンパイル時に圧縮されたファイルを生成することができます．

なお，Brotli と gzip はアルゴリズムが異なることからファイルサイズ以外にも違いが存在し，単純に Brotli に置き換えれば必ず速くなるというわけではありません．そうした詳細については Real-World Effectiveness of Brotli – CSS Wizardry – Web Performance Optimisation が詳しいです．

静的ファイルの配信を最適化
静的ファイルを配信する場合には Express よりも Nginx などのリバースプロキシを使った方がパフォーマンスがよくなります．Express は routing や動的に内容が変化しうる HTML の配信には便利ですが，JS, CSS, 画像といった静的ファイルはリバースプロキシへ移動するとよいです．

もし Nginx の導入が大変という場合でも，Fastify のようなより速いフレームワークに変更することで速度の向上が見込めます．

なお，次項で紹介する CDN を間に挟む場合には Express による影響が抑えられる可能性があります．

キャッシュの設定
開発者が設定できるキャッシュには3段階あります: サーバーサイドキャッシュ，エッジキャッシュ，クライアントサイドキャッシュです．キャッシュは負荷対策の側面とパフォーマンス改善の側面がありますが，ここではパフォーマンスに注目します．

サーバーサイドキャッシュは Redis や lru-cache などの in-memory なキャッシュです．あみぶろでも API が返却した内容とページで表示される内容に差異がない というレギュレーションに抵触しないよう注意が必要ではあるものの，API のリクエスト結果などをキャッシュすることでその分の通信時間を削減できます．

エッジキャッシュは Fastly，CloudCDN，CloudFront などの CDN (Contents Delivery Network) サーバーにおけるキャッシュです．サーバーサイドで用意しているマシンのスペックに依存せず高速に動作すること，ユーザーの地理的位置に応じて最も近いサーバー（エッジノードと呼ばれます）から配信することなどからパフォーマンスの向上が見込めます．

クライアントサイドキャッシュは前述の Service Worker のほか，静的ファイルについては昔から存在する Cache-Control ヘッダーを通じても制御することができます．

なお，どのキャッシュにも言えることですがいわゆる「コールドスタート」問題には注意する必要があります．初回動作時にはキャッシュが存在しないため，速度向上には寄与しません．キャッシュの設定をしたら，計測の前に可能な限りキャッシュをためておく必要があります．同じ理由で，本競技では「再来訪」が発生しないことからクライアントサイドキャッシュは残念ながら効果のない施策となります．

サーバーロケーション・サーバースペックの選定
競技の際デフォルトで提供していたサーバー (Heroku Review Apps) は US のものでした．この場合，地理的に遠い場所（例えば日本）からアクセスするとその分通信に時間がかかることになり，応答時間 (より厳密には，Time to First Byte (TTFB) = リクエストを開始してから始めの 1byte がクライアントへ到着するまでの時間) が遅くなります．

もしアクセスの大半が特定の場所から行われるのであれば，サーバーをその近くに配置することで TTFB を改善できます．何らかの要因でサーバーのリロケーションが難しい場合や，世界中からアクセスがある場合などにはキャッシュの項でも説明した CDN の利用を検討するとよいでしょう．

また，TTFB はサーバー内の処理にかかる時間にも影響を受けるため，単純にサーバーのスペックを上げることでも改善されることがあります．

HTTP/2 による配信
あみぶろの BFF は HTTP/1.1 を使っています．BFF から配信しているリソースの数が多くないためさほど効果が得られない可能性もありますが，HTTP/2 へ移行することは一般にパフォーマンスを向上させます．

歴史を含めた包括的な資料は Introduction to HTTP/2 | Web Fundamentals | Google Developers が参考になります．Express へ導入する場合には node-spdy を使用します．

Server-Side Rendering (SSR)
あみぶろは Cliend-Side Rendering (CSR) をする Single Page Application (SPA) として実装されていますが，CSR は実装によっては First Paint や Time to Interactive が遅くなることがあります．

SSR はこのうち First Paint を向上させます．あみぶろはページ内にユーザー認証情報を含まないため，一般的な Web アプリケーションと比べると比較的容易に SSR が実現できるはずです．生成した HTML をサーバーサイドやエッジノードでキャッシュすることでさらに高速な応答が可能になります．

CSR と SSR 以外の手法も含めると，レンダリングをどこで実行するのかによって Web サイトは大まかに 5 種類に分けられます．Rendering on the Web | Google Developers にて各手法の利点と欠点が網羅的に比較されており，特に最後の表は大変参考になります．

実際に SSR を実装する場合には以下のリソースが参考になります:

ReactDOMServer – React
Server Rendering | Redux
react-router/server-rendering.md at master · ReactTraining/react-router
余談: サポートブラウザを広げる必要がある場合には
競技では Chrome の最新版で動作することのみがレギュレーションであったため，パフォーマンス改善も思い切った対応ができる部分がありました．実際のサービスではサポート範囲をより広く取るため，最新技術に対応していないブラウザのことを考慮に入れる必要が出てくることもあります．この項ではこれまで解説してきた技術の中から，未対応ブラウザに配慮した実装を行う場合をいくつか紹介します．

Differential Serving
トランスパイルや polyfill の使用は一般にバンドルサイズを増加させますが，必要なのは一部の未対応ブラウザだけで，モダンブラウザでは余計なサイズ増加になってしまっているというケースがあります．

そうしたケースに対して，対応しているブラウザに対しては変換しないスリムなコードを配信し，未対応ブラウザに対してのみ必要なトランスパイルや polyfill を施したコードを提供するという技術が考案されており differential serving と呼ばれます．Polyfills の項で紹介した polyfill.io もその一つです．

Differential Serving については筆者の個人ブログに実装方法などをまとめた記事があるため，そちらを参照いただけると嬉しいです:

https://nodaguti.hatenablog.com/entry/2020/04/18/184251
WebP と picture 要素
WebP は確かに画像サイズを小さくすることのできるフォーマットですが，執筆時点では Safari と IE が対応していません．そのため，未対応ブラウザに対しては JPEG や PNG へフォールバックする必要があります．

ブラウザが WebP に対応しているかどうかを調べる一つの方法として Accept ヘッダーをみて image/webp が存在しているかどうかを確かめるというものがありますが，画像ではないリクエストに image/webp を含めるのは本来仕様違反であって，この挙動に依存するのは望ましくありません．

より確実な方法としては <picture> 要素を使う方法があります．これは，以下のように <source> を複数用意することで，ブラウザ自身が読み込み可能なフォーマットを判断して読み込むという挙動に基づくものです．

<picture>
  <source type="image/webp" srcset="foo.webp">
  <source type="image/jpeg" srcset="foo.jpg">
  <img src="foo.jpg" alt="">
</picture>
gzip と Brotli
同じようなフォーマット問題は Brotli にも存在します．こちらは比較的幅広いブラウザでサポートされていますが，古い iOS/Android や IE をサポート対象とする場合には対応が必要です．

Brotli サポートの判別方法は， Accept-Encoding に br という文字列が含まれるかどうかです．実はこのヘッダーの値はブラウザーによってかなり異なるのですが，RFC 7231 に書かれているとおり指定の順序は優先順を表すわけではないので，自前で実装する場合には念頭に入れておきましょう．

パフォーマンスボトルネックを探すには
ここまであみぶろの改善ポイントについて説明してきましたが，あくまで出題者視点からの説明に過ぎませんでした．参加者視点，すなわち実際に未知の Web サイトのパフォーマンスを改善しなければならなくなった場合には，どのようなアプローチがあるのでしょうか．この節では調査に役立つ様々なツールを紹介します．なお，各ツールの使い方についてはそれぞれのリンク先を参照してください．

まずページ全体のパフォーマンスを概観するにはやはり Lighthouse が便利です．Chrome の DevTools や PageSpeed Insights，web.dev から手軽に実行でき，Audits の Opportunities で改善のヒントも得ることができます．

リソースのサイズや読み込みの状況を調査したい場合には DevTools の Network Monitor が便利です (Chrome, Firefox，Safari)．Waterfall 形式で通信内容の詳細や読み込みのタイミングなどを調査することができます．Chrome の DevTools には Code Coverage を調査する機能が搭載されており，未使用の JS/CSS がどれくらいあるのかを数値で示すとともに，行ごとに使用状況を色分けして表示してくれます．

また，JS のバンドルがどのようなモジュールから構成されているのかを調べるには各 bundler のプラグイン (webpack-bundle-analyzer，parcel-plugin-bundle-visualiser，rollup-plugin-visualizer など) が使えるほか，npm ls, Bundlephobia などもバンドルサイズ削減に役立ちます．

ランタイムのパフォーマンス調査には DevTools の Performance Tool が使えます (Chrome，Firefox，Safari)．また，代表的なフレームワークについては個別に特化した DevTools が提供されており，そこでより詳細な処理の様子を観察することもできます (React DevTools，Vue DevTools，Angular Augury など)．

さらに学びたい方へ
First Contentful Paint などの各メトリクスの詳細，RAIL モデル，Performance Budget など Web のパフォーマンス関連で書ききれなかったことはまだまだたくさんあります．また，本稿では主に「どれだけ早くページを表示できるか？」という時間挙動（応答速度）に絞って書いてきましたが，「パフォーマンス」という用語はメモリ使用量などのリソース使用率やスループットなどのキャパシティも含むより広い概念でもあります．もっと詳しく知りたい！と思った方は，以下のリソースにあたってみてください．

オンラインリソース
W3C: Performance and Tuning - Roadmap of Web Applications on Mobile
MDN: Web Performance | MDN
Google: web.dev: Fast load times
Google: Performance | Web Fundamentals | Google Developers
書籍
超速！ Webページ速度改善ガイド ─ 使いやすさは「速さ」から始まる
The Art of Application Performance Testing: From Strategy to Tools
また，パフォーマンス界隈で有名な方のブログも参考になります．例えば:

Addy Osmani
Jason Miller
Ilya Grigorik
 Pages 2
Find a Page…
Home
Web Speed Hackathon Online 出題のねらいと解説
Clone this wiki locally
https://github.com/CyberAgentHack/web-speed-hackathon-online.wiki.git
© 2021 GitHub, Inc.
Terms
Privacy
Security
Status
Help
Contact GitHub
Pricing
API
Training
Blog
About
