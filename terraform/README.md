# Terraform

PLATEAU VIEW 2.0（CMS・エディタ・ビューワ）を構築するためのTerraform用ファイルです。システム構築手順は「実証環境構築マニュアル」も併せて参照してください。

PLATEAU VIEW 2.0のホスティングはGoogle Cloud Platform（GCP）のみ対応しています。AWSやオンプレミスのみでのホスティングはできません。

## セットアップ手順

### GCPのセットアップ

プロジェクトを作成

### MongoDB Atlas のセットアップ

データベースクラスタを作成して接続文字列を取得する。**必ず読取/書込権限を有するDBユーザーの作成・IPアドレスの許可（全IP許可）を忘れずに。**

```bash
export REEARTH_DB=""
```

### Auth0 のセットアップ

Auth0テナントを作成した後、[公式のQuick Start](https://github.com/auth0/terraform-provider-auth0/blob/main/docs/guides/quickstart.md)を参考に、アプリケーションをセットアップ。

```bash
export AUTH0_CLIENT_SECRET=""
```

※2度目以降の `terraform apply` でもこの変数の設定が必要です。

### コマンドラインツールのインストール

公式ドキュメントに従ってインストール。

 - gcloud
 - terraform

### gcloud のセットアップ

```bash
# GCP Project ID
export PROJECT_ID=""
# 使いたいドメイン 例: plateauview.example.com
export DOMAIN=""
# 20文字以内・半角英数ハイフンで自由に決めて良い。例: plateauview-test
export SERVICE_PREFIX=""

gcloud components update
gcloud config configurations create ${SERVICE_PREFIX}
gcloud config set project ${PROJECT_ID}
gcloud auth login

# GCPのAPIの有効化。中には完了まで少し時間を要するものもある。
gcloud services enable certificatemanager.googleapis.com
gcloud services enable secretmanager.googleapis.com
gcloud services enable cloudbuild.googleapis.com
gcloud services enable cloudresourcemanager.googleapis.com
gcloud services enable cloudtasks.googleapis.com
gcloud services enable compute.googleapis.com
gcloud services enable dns.googleapis.com
gcloud services enable iam.googleapis.com
gcloud services enable run.googleapis.com
gcloud services enable sts.googleapis.com

gcloud config set compute/region asia-northeast1
gcloud auth application-default login
```

### Cloud DNS のセットアップ

descriptionがないとエラーになるので注意。

```bash
gcloud dns managed-zones create ${SERVICE_PREFIX} --dns-name ${DOMAIN} --description "${SERVICE_PREFIX}"
gcloud dns record-sets list --zone ${ZONE_NAME}
```

以下のような出力が得られる。

```
NAME                           TYPE  TTL    DATA
*********  NS    21600  ns-cloud-a1.googledomains.com.,ns-cloud-a2.googledomains.com.,ns-cloud-a3.googledomains.com.,ns-cloud-a4.googledomains.com.
*********  SOA   21600  ns-cloud-a1.googledomains.com. cloud-dns-hostmaster.google.com. 1 21600 3600 259200 300
```

今回セットアップしたいドメインを、Cloud DNSでホスティングできるように、各種レジストラの設定を変更する。具体的には、出力のNSレコードのDATAの部分を使用して「NSレコード」を設定することで、ネームサーバーを変更する。手順は各種レジストラの手順を参照。

### 設定ファイルの準備

今回設定したい環境の設定ファイル(tfvars)を準備する。実行環境費に合わせて必要な情報をセットアップを行う。

[example.tfvars](./env/example.tfvars) を編集して必要な設定を追記する。（別名としてコピーして編集しても良い）

### Terraform の backend を作成

```bash
gcloud storage buckets create gs://${SERVICE_PREFIX}-terraform-tfstate
```

[terraform.tf](terraform.tf) の `backend` の `bucket` を変更する

```diff
  backend "gcs" {
-    bucket = ""
+    bucket = "${SERVICE_PREFIXで指定した値を入れる}-terraform-tfstate"
  }
```

## Terraform実行後手順

```bash
terraform init
```

```bash
terraform apply -var-file=env/example.tfvars
```

途中でyesを入力して実行を進めること。しばらく時間を要するが、以下のような出力が得られれば成功。

```
Apply complete! Resources: * added, * changed, * destroyed.

Outputs:

plateauview_cms_url = "********"
plateauview_cms_webhook_secret = "********"
plateauview_cms_webhook_url = "********"
plateauview_reearth_url = "********"
plateauview_sdk_token = "********"
plateauview_sidebar_token = "********"
plateauview_sidecar_url = "********"
```

これらの outputs は後で使う。なおもう一度 outputs を表示したいときは `terraform output` コマンドで表示可能。

- `plateauview_cms_url`: CMS（Re:Earth CMS）のURL
- `plateauview_cms_webhook_secret`: 下記「CMS インテグレーション設定」で使用
- `plateauview_cms_webhook_url`: 下記「CMS インテグレーション設定」で使用
- `plateauview_reearth_url`: エディタ（Re:Earth）のURL
- `plateauview_sdk_token`: PLATEAU SDK用のトークン。SDKのUIで設定する（詳しくは実証環境構築マニュアルを参照）。
- `plateauview_sidebar_token`: ビューワのサイドバー用のAPIトークン。エディタ上でサイドバーウィジェットの設定から設定する（詳しくは実証環境構築マニュアルを参照）。
- `plateauview_sidecar_url`: サイドカーサーバーのURL。エディタ上でサイドバーウィジェットの設定から設定する（詳しくは実証環境構築マニュアルを参照）。

### シークレットの設定

GCPのシークレットマネージャにシークレットを設定する。

```bash
echo -n "${REEARTH_DB}" | gcloud secrets versions add reearth-api-REEARTH_DB --data-file=-
echo -n "${REEARTH_DB}" | gcloud secrets versions add reearth-cms-REEARTH_CMS_WORKER_DB --data-file=-
echo -n "${REEARTH_DB}" | gcloud secrets versions add reearth-cms-REEARTH_CMS_DB --data-file=-
```

なお、以下は必要に応じて設定する。各 `${...}` は変更して実行すること。設定しなくても次に進むことは可能。

```bash
# FMEのトークン
echo -n "${REEARTH_PLATEAUVIEW_FME_TOKEN}" | gcloud secrets versions add reearth-cms-REEARTH_PLATEAUVIEW_FME_TOKEN --data-file=-
# G空間情報センターのAPIトークン
echo -n "${REEARTH_PLATEAUVIEW_CKAN_TOKEN}" | gcloud secrets versions add reearth-cms-REEARTH_PLATEAUVIEW_CKAN_TOKEN --data-file=-
# SendGridのAPIキー
echo -n "${REEARTH_PLATEAUVIEW_SENDGRID_APIKEY}" | gcloud secrets versions add reearth-cms-REEARTH_PLATEAUVIEW_SENDGRID_APIKEY --data-file=-
# マーケットプレースのシークレットキー
echo -n "${REEARTH_MARKETPLACE_SECRET}" | gcloud secrets versions add reearth-api-REEARTH_MARKETPLACE_SECRET --data-file=-
```

### Cloud Run のデプロイ

4つの Cloud Run サービスをデプロイする。

```bash
gcloud run deploy reearth-api \
  --image eukarya/plateauview2-reearth:latest \
  --region asia-northeast1 \
  --platform managed \
  --quiet
```

```bash
gcloud run deploy reearth-cms-api \
  --image eukarya/plateauview2-reearth-cms:latest \
  --region asia-northeast1 \
  --platform managed \
  --quiet
```

```bash
gcloud run deploy reearth-cms-worker \
  --image eukarya/plateauview2-reearth-cms-worker:latest \
  --region asia-northeast1 \
  --platform managed \
  --quiet
```

```bash
gcloud run deploy plateauview-api \
  --image eukarya/plateauview2-sidecar:latest \
  --region asia-northeast1 \
  --platform managed \
  --quiet
```

### DNS・ロードバランサ・証明書のデプロイ完了まで待機

```bash
curl https://api.${DOMAIN}/ping
```

を繰り返し試行し `"pong"` が返ってくるまで待つ。

### Auth0 ユーザー作成


先ほど作成したAuth0テナントにてユーザーを作成する。メールアドレスの認証を忘れずに。

**必ず上記ステップでデプロイ完了を確認してから、Auth0のユーザーを作成すること。そうでないと正常にRe:EarthやCMSにログインできなくなる。**

### CMS インテグレーション設定

Terraformのoutputsの `plateauview_cms_url` のURL（`https://reearth.${DOMAIN}`）から、CMSにログインする。

ログイン後、ワークスペース・Myインテグレーションを作成する。

次に、インテグレーション内に以下の通り webhook を作成する。作成後、有効化を忘れないこと。

- URL: terraform outputs の plateauview_cms_webhook_url
- シークレット: terraform outputs の plateauview_cms_webhook_secret
- イベント: 全てのチェックボックスにチェックを入れる。

作成後、作成したワークスペースに作成したインテグレーションを追加し、オーナー権限に変更する。

先ほど作成したインテグレーションの詳細画面でインテグレーショントークンをコピーし、以下の `${REEARTH_PLATEAUVIEW_CMS_TOKEN}` に貼り付けて以下のコマンドを実行する。

```bash
echo -n "${REEARTH_PLATEAUVIEW_CMS_TOKEN}" | gcloud secrets versions add reearth-cms-REEARTH_PLATEAUVIEW_CMS_TOKEN --data-file=-
```

環境変数の変更を適用するため、もう一度 Cloud Run をデプロイする。

```bash
gcloud run deploy plateauview-api \
  --image eukarya/plateauview2-sidecar:latest \
  --region asia-northeast1 \
  --platform managed \
  --quiet
```

### 完了

以下のアプリケーションにログインし、正常に使用できることを確認する。 `${DOMAIN}` はドメイン。

- Re:Earth: Terraformのoutputsの `plateauview_reearth_url` の値（`https://reearth.${DOMAIN}`）
- CMS: Terraformのoutputsの `plateauview_cms_url` の値（`https://cms.${DOMAIN}`）
