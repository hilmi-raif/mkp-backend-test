begin;

create table pengguna (
    id_pengguna bigint generated always as identity primary key,
    email text not null unique,
    password_hash text not null
);

create table cabang (
    id_cabang bigint generated always as identity primary key,
    nama text not null,
    kota text not null,
    unique (kota, nama)
);

create table studio (
    id_studio bigint generated always as identity primary key,
    id_cabang bigint not null references cabang,
    nama text not null,
    unique (id_cabang, nama)
);

create table kursi (
    id_kursi bigint generated always as identity primary key,
    id_studio bigint not null references studio,
    baris text not null,
    nomor integer not null check (nomor > 0),
    unique (id_studio, baris, nomor),
    unique (id_kursi, id_studio)
);

create table film (
    id_film bigint generated always as identity primary key,
    judul text not null
);

create table jadwal_tayang (
    id_jadwal bigint generated always as identity primary key,
    id_studio bigint not null references studio,
    id_film bigint not null references film,
    mulai_at timestamptz not null,
    selesai_at timestamptz not null,
    harga bigint not null check (harga > 0),
    status text not null default 'aktif' check (status in ('aktif', 'penjualan_tutup', 'batal')),
    check (selesai_at > mulai_at),
    unique (id_jadwal, id_studio)
);

create index jadwal_studio_mulai_idx on jadwal_tayang (id_studio, mulai_at);

-- status kursi sekarang untuk tiap jadwal
create table kursi_jadwal (
    id_jadwal bigint not null,
    id_kursi bigint not null,
    id_studio bigint not null,
    id_detail_aktif bigint,
    status text not null default 'tersedia' check (status in ('tersedia', 'ditahan', 'terjual')),
    primary key (id_jadwal, id_kursi),
    foreign key (id_jadwal, id_studio) references jadwal_tayang (id_jadwal, id_studio),
    foreign key (id_kursi, id_studio) references kursi (id_kursi, id_studio),
    check ((status = 'tersedia') = (id_detail_aktif is null))
);

create table pesanan (
    id_pesanan bigint generated always as identity primary key,
    id_pengguna bigint not null references pengguna,
    id_jadwal bigint not null references jadwal_tayang,
    dibuat_at timestamptz not null default now(),
    batas_bayar_at timestamptz not null,
    status text not null default 'menunggu_pembayaran'
        check (status in ('menunggu_pembayaran', 'dibayar', 'kedaluwarsa', 'batal')),
    unique (id_pesanan, id_jadwal)
);

-- riwayat kursi tetap ada walaupun pemegangnya ganti
create table detail_pesanan (
    id_detail bigint generated always as identity primary key,
    id_pesanan bigint not null,
    id_jadwal bigint not null,
    id_kursi bigint not null,
    harga bigint not null check (harga > 0),
    foreign key (id_pesanan, id_jadwal) references pesanan (id_pesanan, id_jadwal),
    foreign key (id_jadwal, id_kursi) references kursi_jadwal (id_jadwal, id_kursi),
    unique (id_pesanan, id_kursi),
    unique (id_detail, id_jadwal, id_kursi)
);

alter table kursi_jadwal add foreign key (id_detail_aktif, id_jadwal, id_kursi)
    references detail_pesanan (id_detail, id_jadwal, id_kursi);

create table pembayaran (
    id_pembayaran bigint generated always as identity primary key,
    id_pesanan bigint not null references pesanan,
    referensi_gateway text unique,
    jumlah bigint not null check (jumlah > 0),
    status text not null default 'menunggu_konfirmasi'
        check (status in ('menunggu_konfirmasi', 'dana_diterima', 'gagal', 'cek_manual')),
    dibuat_at timestamptz not null default now()
);

create table tiket (
    id_tiket bigint generated always as identity primary key,
    id_detail bigint not null unique,
    id_jadwal bigint not null,
    id_kursi bigint not null,
    kode text not null unique,
    status text not null default 'aktif' check (status in ('aktif', 'batal')),
    terbit_at timestamptz not null default now(),
    foreign key (id_detail, id_jadwal, id_kursi) references detail_pesanan (id_detail, id_jadwal, id_kursi)
);

-- tiket lama yang batal tidak menghalangi penjualan ulang
create unique index tiket_aktif_per_kursi_idx on tiket (id_jadwal, id_kursi) where status = 'aktif';

create table refund (
    id_refund bigint generated always as identity primary key,
    id_pembayaran bigint not null references pembayaran,
    kunci_refund text not null unique,
    jumlah bigint not null check (jumlah > 0),
    status text not null default 'menunggu' check (status in ('menunggu', 'diajukan', 'berhasil', 'cek_manual')),
    dibuat_at timestamptz not null default now()
);

commit;
